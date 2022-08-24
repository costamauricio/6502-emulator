package cpu6502

import (
    "6502_emulator/bus"
)

type Flag = byte

const (
    FLAG_C Flag = 1<<0 // Carry bit
    FLAG_Z Flag = 1<<1 // Zero
    FLAG_I Flag = 1<<2 // Disable interrupts
    FLAG_D Flag = 1<<3 // Decimal mode
    FLAG_B Flag = 1<<4 // Break Command
    FLAG_U Flag = 1<<5 // unused
    FLAG_V Flag = 1<<6 // Overflow
    FLAG_N Flag = 1<<7 // Negative
)

type CPU struct {
    A byte // Accumulator
    X byte // X Register
    Y byte // Y Register
    S byte // Stack Pointer
    PC uint16 // Program Counter
    Status byte

    absolute_address uint16
    bus *bus.Bus
}

func (cpu *CPU) GetFlag(flag Flag) bool {
    return (cpu.Status & flag) > 0
}

func (cpu *CPU) SetFlag(flag Flag, value bool) {
    if value {
        cpu.Status = cpu.Status | flag
        return
    }

    cpu.Status = cpu.Status &^ flag
}

func (cpu *CPU) Clock() {

}

func (cpu *CPU) write(address uint16, data byte) {
    cpu.bus.Write(address, data)
}

func (cpu *CPU) read(address uint16) byte {
    return cpu.bus.Read(address)
}
