package main

import (
	"emulator/pkg/bus"
	"emulator/pkg/cpu6502"
	"emulator/internal/visualizer"
	"log"
)

func main() {
	dataBus := bus.Bus{}
	dataBus.LoadRamFromString("A9 0A 69 02 AA 86 01 E9 02 D0 FC 00", 0x8000)
	dataBus.LoadRamFromString("00 80", 0xFFFC)

	cpu := cpu6502.New(&dataBus)

	log.Print("CPU: ", cpu)
	visualizer := visualizer.Visualizer{Cpu: cpu, Bus: &dataBus}
	visualizer.Run(0x8000)
}
