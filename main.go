package main

import (
	"6502_emulator/bus"
	"6502_emulator/cpu6502"
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
    cpu *cpu6502.CPU
    dataBus bus.Bus
    font *ttf.Font
    window *sdl.Window
    whiteColor sdl.Color
    redColor sdl.Color
)

func init() {
    whiteColor = sdl.Color{255, 255, 255, 0}
    redColor = sdl.Color{255, 0, 0, 0}
}

func main() {
    var err error

    if err = ttf.Init(); err != nil {
	    panic(err)
	}
	defer ttf.Quit()

    if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

    if font, err = ttf.OpenFont("assets/pixel-font.ttf", 16); err != nil {
        panic(err)
    }
    defer font.Close()

	dataBus = bus.Bus{}
	cpu = cpu6502.New(&dataBus)

    dataBus.LoadRamFromString("0E FE 82 A6 D6", 0x8000)
    dataBus.LoadRamFromString("00 80", 0xFFFC)

    log.Print("CPU: ", cpu)

    window, err = sdl.CreateWindow(
        "6502 Emulator",
        sdl.WINDOWPOS_UNDEFINED,
        sdl.WINDOWPOS_UNDEFINED,
		800,
        600,
        sdl.WINDOW_SHOWN | sdl.WINDOW_RESIZABLE)

	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	running := true
	for running {
        drawRam(2, 2, 0x0000)
        drawRam(2, 280, 0x8000)
        drawCpu()
        drawCommands()
	    window.UpdateSurface()

        event := sdl.PollEvent()

        if event == nil {
            continue
        }

        switch event := event.(type) {
        case *sdl.QuitEvent:
            println("Quit")
            running = false
        case *sdl.TextInputEvent:
            println(event.GetText())
        }
	}
}

func drawText(text string, x int32, y int32, color *sdl.Color) {
    surface, err := window.GetSurface()
    if err != nil {
        panic(err)
    }

    if color == nil {
        color = &whiteColor
    }

    textSurface, _ := font.RenderUTF8Solid(text, *color)
    defer textSurface.Free()

    textSurface.Blit(nil, surface, &sdl.Rect{X: x, Y: y, W: 0, H: 0})
}

func drawRam(x int32, y int32, offset uint16) {
    for rows := uint16(0); rows < 16; rows++ {
        var line = ""
        line += fmt.Sprintf("$%04X", rows*16 + offset) + ": "
        for index := uint16(0); index < 16; index++ {
            line += fmt.Sprintf("%02X", dataBus.Read(rows*16 + index + offset)) + " "
        }

        drawText(line, x, int32(rows)*16 + y, nil)
    }
}

func drawCpu() {
    var x, y int32 = 500, 2
    drawText("STATUS: ", x, y, nil)

    testFlag := func(tested cpu6502.Flag) *sdl.Color {
        if cpu.GetFlag(tested) > 0 {
            return &redColor
        }

        return nil
    }

    x += 64
    drawText("N", x, y, testFlag(cpu6502.FLAG_N))
    drawText("V", x+16, y, testFlag(cpu6502.FLAG_V))
    drawText("-", x+32, y, testFlag(cpu6502.FLAG_U))
    drawText("B", x+48, y, testFlag(cpu6502.FLAG_B))
    drawText("D", x+64, y, testFlag(cpu6502.FLAG_D))
    drawText("I", x+80, y, testFlag(cpu6502.FLAG_I))
    drawText("Z", x+96, y, testFlag(cpu6502.FLAG_Z))
    drawText("C", x+112, y, testFlag(cpu6502.FLAG_C))
    x -= 64

    drawText(fmt.Sprintf("PC: $%04X", cpu.PC), x, y+16, nil)
    drawText(fmt.Sprintf("A: $%02X  [%d]", cpu.A, cpu.A), x, y+32, nil)
    drawText(fmt.Sprintf("X: $%02X  [%d]", cpu.X, cpu.X), x, y+48, nil)
    drawText(fmt.Sprintf("Y: $%02X  [%d]", cpu.Y, cpu.Y), x, y+64, nil)
    drawText(fmt.Sprintf("P: $%02X", cpu.S), x, y+80, nil)
}

func drawCommands() {
    var x, y int32 = 10, 550

    drawText("SPACE = Step Instruction", x, y, nil);
    drawText("R = Reset", x+232, y, nil);
    drawText("I = IRQ", x+344, y, nil);
    drawText("N = NMI", x+440, y, nil);
}
