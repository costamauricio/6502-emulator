package debugger

import (
	"6502_emulator/bus"
	"6502_emulator/cpu6502"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	white string = "WHITE"
	black string = "BLACK"
	red   string = "RED"
	blue  string = "BLUE"
)

type Visualizer struct {
	Cpu *cpu6502.CPU
	Bus *bus.Bus

	font     *ttf.Font
	renderer *sdl.Renderer

	colors map[string]*sdl.Color
}

// Runs the visualizer
// initMemory uint16 - The program memory starts at
func (v *Visualizer) Run(initMemory uint16) error {
	v.colors = make(map[string]*sdl.Color)

	v.colors[white] = &sdl.Color{255, 255, 255, 0}
	v.colors[black] = &sdl.Color{0, 0, 0, 0}
	v.colors[red] = &sdl.Color{255, 0, 0, 0}
	v.colors[blue] = &sdl.Color{29, 76, 186, 0}

	var err error

	if err = ttf.Init(); err != nil {
		return err
	}
	defer ttf.Quit()

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}
	defer sdl.Quit()

	if v.font, err = ttf.OpenFont("assets/pixel-font.ttf", 16); err != nil {
		return err
	}
	defer v.font.Close()

	window, err := sdl.CreateWindow(
		"6502 Emulator",
		sdl.WINDOWPOS_CENTERED_MASK,
		sdl.WINDOWPOS_CENTERED_MASK,
		800,
		600,
		sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)

	if err != nil {
		return err
	}
	defer window.Destroy()

	if v.renderer, err = sdl.CreateRenderer(window, -1, 0); err != nil {
		return err
	}
	defer v.renderer.Destroy()

	parsedCode, codeOrder := v.Cpu.DisassembleInstructions(0x0000, 0xFFFF)
	v.Cpu.Reset()

	running := true
	for running {
		v.renderer.SetDrawColor(77, 77, 77, 0)
		v.renderer.Clear()

		v.drawRam("Zero Page", 20, 20, 16, 0x0000)
		v.drawRam("Stack", 20, 308, 4, 0x01BF)
		v.drawRam("Program", 20, 404, 8, initMemory)
		v.drawCpu()
		v.drawInstructions(parsedCode, codeOrder)
		v.drawCommands()

		v.renderer.Present()

		event := sdl.PollEvent()

		if event == nil {
			continue
		}

		switch event := event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			running = false
		case *sdl.TextInputEvent:
			switch event.GetText() {
			case " ":
				v.stepInstruction()
			case "r", "R":
				v.Cpu.Reset()
			case "i", "I":
				v.Cpu.InterruptRequest()
			case "n", "N":
				v.Cpu.NonMaskableInterrupt()
			}
		}
	}

	return nil
}

func (v *Visualizer) stepInstruction() {
	for {
		v.Cpu.Tick()
		if v.Cpu.InstructionCompleted() {
			break
		}
	}
}

func (v *Visualizer) drawText(text string, x int32, y int32, color *sdl.Color) {
	if color == nil {
		color = v.colors[white]
	}

	//textSurface, _ := v.font.RenderUTF8Shaded(text, *color, *v.colors[black])
	textSurface, _ := v.font.RenderUTF8Blended(text, *color)
	defer textSurface.Free()

	texture, _ := v.renderer.CreateTextureFromSurface(textSurface)
	defer texture.Destroy()

	v.renderer.Copy(texture, nil, &sdl.Rect{X: x, Y: y, W: textSurface.W, H: textSurface.H})
}

func (v *Visualizer) drawShadedText(text string, x int32, y int32, color *sdl.Color, fg *sdl.Color) {
	if color == nil {
		color = v.colors[white]
	}

	textSurface, _ := v.font.RenderUTF8Shaded(text, *color, *fg)
	defer textSurface.Free()

	texture, _ := v.renderer.CreateTextureFromSurface(textSurface)
	defer texture.Destroy()

	v.renderer.Copy(texture, nil, &sdl.Rect{X: x, Y: y, W: textSurface.W, H: textSurface.H})
}

func (v *Visualizer) drawBox(text string, x int32, y int32, l int32, h int32) {
	v.renderer.DrawRect(&sdl.Rect{x, y, l, h})

	if text != "" {
		v.drawShadedText(text, x+16, y-8, nil, &sdl.Color{77, 77, 77, 0})
	}
}

func (v *Visualizer) drawRam(label string, x int32, y int32, lines int, offset uint16) {
	v.renderer.SetDrawColor(0, 0, 0, 0)
	v.drawBox(label, x-10, y-10, 432+20, int32(16*lines+20))
	for rows := uint16(0); rows < uint16(lines); rows++ {
		var line = ""
		line += fmt.Sprintf("$%04X", rows*16+offset) + ": "
		for index := uint16(0); index < 16; index++ {
			line += fmt.Sprintf("%02X", v.Bus.Read(rows*16+index+offset)) + " "
		}

		v.drawText(line, x, int32(rows)*16+y, nil)
	}
}

func (v *Visualizer) drawCpu() {
	var x, y int32 = 480, 2
	v.drawText("STATUS: ", x, y, nil)

	testFlag := func(tested cpu6502.Flag) *sdl.Color {
		if v.Cpu.GetFlag(tested) > 0 {
			return v.colors[red]
		}

		return nil
	}

	x += 64
	v.drawText("N", x, y, testFlag(cpu6502.FLAG_N))
	v.drawText("V", x+16, y, testFlag(cpu6502.FLAG_V))
	v.drawText("-", x+32, y, testFlag(cpu6502.FLAG_U))
	v.drawText("B", x+48, y, testFlag(cpu6502.FLAG_B))
	v.drawText("D", x+64, y, testFlag(cpu6502.FLAG_D))
	v.drawText("I", x+80, y, testFlag(cpu6502.FLAG_I))
	v.drawText("Z", x+96, y, testFlag(cpu6502.FLAG_Z))
	v.drawText("C", x+112, y, testFlag(cpu6502.FLAG_C))
	x -= 64

	v.drawText(fmt.Sprintf("PC: $%04X", v.Cpu.PC), x, y+16, nil)
	v.drawText(fmt.Sprintf("A: $%02X  [%d]", v.Cpu.A, v.Cpu.A), x, y+32, nil)
	v.drawText(fmt.Sprintf("X: $%02X  [%d]", v.Cpu.X, v.Cpu.X), x, y+48, nil)
	v.drawText(fmt.Sprintf("Y: $%02X  [%d]", v.Cpu.Y, v.Cpu.Y), x, y+64, nil)
	v.drawText(fmt.Sprintf("P: $%02X", v.Cpu.S), x, y+80, nil)
}

func (v *Visualizer) drawInstructions(instructions map[uint16]string, order []uint16) {
	var x, y int32 = 480, 130

	var currentIndex int

	getColor := func(current bool) *sdl.Color {
		if current {
			return &sdl.Color{0, 255, 0, 0}
		}

		return nil
	}

	for index := range order {
		if order[index] == v.Cpu.PC {
			currentIndex = index
			break
		}
	}

	topCut, bottomCut := currentIndex-13, currentIndex+13

	if topCut < 0 {
		topCut = 0
	}

	if bottomCut > len(order) {
		bottomCut = len(order)
	}

	for index, value := range order[topCut:bottomCut] {
		v.drawText(instructions[value], x, y+(int32(index)*16), getColor(value == v.Cpu.PC))
	}
}

func (v *Visualizer) drawCommands() {
	var x, y int32 = 10, 550

	v.drawText("SPACE = Step Instruction", x, y, nil)
	v.drawText("R = Reset", x+232, y, nil)
	v.drawText("I = IRQ", x+344, y, nil)
	v.drawText("N = NMI", x+440, y, nil)
}
