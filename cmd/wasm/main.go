package main

import (
	"github.com/costamauricio/6502-emulator/pkg/bus"
	"github.com/costamauricio/6502-emulator/pkg/cpu6502"
	"syscall/js"
)

var (
	dataBus *bus.Bus
	cpu     *cpu6502.CPU
)

func main() {
	dataBus = &bus.Bus{}
	cpu = cpu6502.New(dataBus)

	js.Global().Set("getRegisters", js.FuncOf(getRegisters))
	js.Global().Set("getFlags", js.FuncOf(getFlags))

	<-make(chan bool)
}

//export reset
func reset() {
    cpu.Reset()
}

//export stepInstruction
func stepInstruction() {
	for {
		cpu.Tick()
		if cpu.InstructionCompleted() {
			break
		}
	}
}

func getRegisters(this js.Value, args []js.Value) interface{} {
	return map[string]interface{}{
		"a":               cpu.A,
		"y":               cpu.Y,
		"x":               cpu.X,
		"status":          cpu.Status,
		"stack_pointer":   cpu.S,
		"program_counter": cpu.PC,
	}
}

func getFlags(this js.Value, args []js.Value) interface{} {
	return map[string]interface{}{
		"n": cpu.GetFlag(cpu6502.FLAG_N) > 0,
		"v": cpu.GetFlag(cpu6502.FLAG_V) > 0,
		"u": cpu.GetFlag(cpu6502.FLAG_U) > 0,
		"b": cpu.GetFlag(cpu6502.FLAG_B) > 0,
		"d": cpu.GetFlag(cpu6502.FLAG_D) > 0,
		"i": cpu.GetFlag(cpu6502.FLAG_I) > 0,
		"z": cpu.GetFlag(cpu6502.FLAG_Z) > 0,
		"c": cpu.GetFlag(cpu6502.FLAG_C) > 0,
	}
}
