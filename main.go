package main

import (
	"6502_emulator/bus"
	"6502_emulator/cpu6502"
    "log"
)

func main() {
	dataBus := bus.Bus{}
	cpu := cpu6502.New(&dataBus)

    log.Print("CPU: ", cpu)
}
