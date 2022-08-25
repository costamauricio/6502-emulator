package main

import (
	"6502_emulator/bus"
	"6502_emulator/cpu6502"
)

func main() {
	cpu := cpu6502.CPU{}
	dataBus := bus.Bus{}

	cpu.AttachToBus(&dataBus)
}
