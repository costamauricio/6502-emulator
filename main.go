package main

import (
	"6502_emulator/bus"
	"6502_emulator/cpu6502"
	"6502_emulator/debugger"
	"log"
)

func main() {
	dataBus := bus.Bus{}
	dataBus.LoadRamFromString("A9 0A 69 02 AA 86 01 E9 02 D0 FC 00", 0x8000)
	dataBus.LoadRamFromString("00 80", 0xFFFC)

	cpu := cpu6502.New(&dataBus)

	log.Print("CPU: ", cpu)
	visualizer := debugger.Visualizer{Cpu: cpu, Bus: &dataBus}
	visualizer.Run(0x8000)
}
