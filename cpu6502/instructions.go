package cpu6502

type Instruction string

const (
    INS_ADC Instruction = "ADC"
    INS_AND Instruction = "AND"
    INS_LDA Instruction = "LDA"
    INS_LDX Instruction = "LDX"
    INS_LDY Instruction = "LDY"
)

func attachInstructions(cpu *CPU) {
    cpu.instructions = make(Instructions)

    cpu.instructions[INS_ADC] = cpu.adc
    cpu.instructions[INS_AND] = cpu.and
    cpu.instructions[INS_LDA] = cpu.lda
    cpu.instructions[INS_LDX] = cpu.ldx
    cpu.instructions[INS_LDY] = cpu.ldy
}

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
func (cpu *CPU) adc(mode AddressingMode) {

}

// And memory with accumulator
func (cpu *CPU) and(mode AddressingMode) {
    data := cpu.loadData(mode)
    cpu.A = cpu.A & data

    cpu.SetFlag(FLAG_Z, cpu.A == 0x00)
    cpu.SetFlag(FLAG_N, cpu.A & 0x80 > 0)
}

// Load accumulator with memory
func (cpu *CPU) lda(mode AddressingMode) {
    data := cpu.loadData(mode)
    cpu.A = data

    cpu.SetFlag(FLAG_Z, cpu.A == 0x00)
    cpu.SetFlag(FLAG_N, cpu.A & 0x80 > 0)
}

// Load index X with memory
func (cpu *CPU) ldx(mode AddressingMode) {
    data := cpu.loadData(mode)
    cpu.X = data

    cpu.SetFlag(FLAG_Z, cpu.X == 0x00)
    cpu.SetFlag(FLAG_N, cpu.X & 0x80 > 0)
}

// Load index Y with memory
func (cpu *CPU) ldy(mode AddressingMode) {
    data := cpu.loadData(mode)
    cpu.Y = data

    cpu.SetFlag(FLAG_Z, cpu.Y == 0x00)
    cpu.SetFlag(FLAG_N, cpu.Y & 0x80 > 0)
}
