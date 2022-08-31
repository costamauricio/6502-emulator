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

type AddressingModes map[AddressingMode]func()uint16
type Instructions map[Instruction]func(AddressingMode)

type CPU struct {
    A byte // Accumulator
    X byte // X Register
    Y byte // Y Register
    S byte // Stack Pointer
    PC uint16 // Program Counter
    Status byte

    bus *bus.Bus
    cicles int // Current instruction cicles

    addressingModes AddressingModes
    instructions Instructions
}

// Initialize a new CPU
func New(bus *bus.Bus) *CPU {
    cpu := CPU{bus: bus}

    attachAddressModes(&cpu)
    attachInstructions(&cpu)

    cpu.Reset()
    return &cpu
}

// Flags
func (cpu *CPU) GetFlag(flag Flag) byte {
    return cpu.Status & flag
}

func (cpu *CPU) SetFlag(flag Flag, value bool) {
    if value {
        cpu.Status = cpu.Status | flag
        return
    }

    cpu.Status = cpu.Status &^ flag
}

// Data Bus
func (cpu *CPU) write(address uint16, data byte) {
    cpu.bus.Write(address, data)
}

func (cpu *CPU) read(address uint16) byte {
    return cpu.bus.Read(address)
}

// Perform a CPU clock cicle
func (cpu *CPU) Tick() {
    if cpu.cicles > 0 {
        cpu.cicles--
        return
    }

    //opcode := cpu.read(cpu.PC)
    //cpu.PC++

    // Continue with the mappings
    cpu.cicles--
}

func (cpu *CPU) Reset() {
    var progAddress uint16 = 0xFFFC;

    low := uint16(cpu.read(progAddress))
    high := uint16(cpu.read(progAddress + 1))

    cpu.PC = (high << 8) | low
    cpu.A = 0x00
    cpu.X = 0x00
    cpu.Y = 0x00
    cpu.S = 0xFD
    cpu.Status = 0x00 | FLAG_U

    // It takes 6 cicles to the CPU to restart
    cpu.cicles = 6
}

func (cpu *CPU) InterruptRequest() {

}

func (cpu *CPU) NonMaskableInterrupt() {

}
