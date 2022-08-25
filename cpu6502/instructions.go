package cpu6502

func (cpu *CPU) loadData(mode AddressingMode) byte {
    if mode == MODE_ACC {
        return cpu.A
    }

    if mode == MODE_IMP {
        return 0
    }

    address := cpu.addressingModes[mode]()
    return cpu.read(address)
}

// Addition
// Add memory to accumulator with carry bit
func adc(cpu *CPU, mode AddressingMode) {

}

// And memory with accumulator
func and(cpu *CPU, mode AddressingMode) {
    data := cpu.loadData(mode)
    cpu.A = cpu.A & data

    cpu.SetFlag(FLAG_Z, cpu.A == 0x00)
    cpu.SetFlag(FLAG_N, cpu.A & 0x80 > 0)
}

// Load accumulator with memory
func lda(cpu *CPU, mode AddressingMode) {
    data := cpu.loadData(mode)
    cpu.A = data

    cpu.SetFlag(FLAG_Z, cpu.A == 0x00)
    cpu.SetFlag(FLAG_N, cpu.A & 0x80 > 0)
}

// Load index X with memory
func ldx(cpu *CPU, mode AddressingMode) {
    data := cpu.loadData(mode)
    cpu.X = data

    cpu.SetFlag(FLAG_Z, cpu.X == 0x00)
    cpu.SetFlag(FLAG_N, cpu.X & 0x80 > 0)
}

// Load index Y with memory
func ldy(cpu *CPU, mode AddressingMode) {
    data := cpu.loadData(mode)
    cpu.Y = data

    cpu.SetFlag(FLAG_Z, cpu.Y == 0x00)
    cpu.SetFlag(FLAG_N, cpu.Y & 0x80 > 0)
}
