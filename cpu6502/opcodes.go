package cpu6502

type opcode struct {
    instruction Instruction
    addressMode AddressingMode
    cicles int
}

var OPCODES map[byte]opcode

func init() {
    OPCODES = make(map[byte]opcode)

    OPCODES[0x69] = opcode{ INS_ADC, MODE_IMM, 2 }
    OPCODES[0x6D] = opcode{ INS_ADC, MODE_ABS, 4 }
    OPCODES[0x65] = opcode{ INS_ADC, MODE_ZP0, 3 }
    OPCODES[0x61] = opcode{ INS_ADC, MODE_INX, 6 }
    OPCODES[0x71] = opcode{ INS_ADC, MODE_INY, 5 }
    OPCODES[0x75] = opcode{ INS_ADC, MODE_ZPX, 4 }
    OPCODES[0x7D] = opcode{ INS_ADC, MODE_ABX, 4 }
    OPCODES[0x79] = opcode{ INS_ADC, MODE_ABY, 4 }

    OPCODES[0x29] = opcode{ INS_AND, MODE_IMM, 2 }
    OPCODES[0x2D] = opcode{ INS_AND, MODE_ABS, 4 }
    OPCODES[0x25] = opcode{ INS_AND, MODE_ZP0, 3 }
    OPCODES[0x21] = opcode{ INS_AND, MODE_INX, 6 }
    OPCODES[0x31] = opcode{ INS_AND, MODE_INY, 5 }
    OPCODES[0x35] = opcode{ INS_AND, MODE_ZPX, 4 }
    OPCODES[0x3D] = opcode{ INS_AND, MODE_ABX, 4 }
    OPCODES[0x39] = opcode{ INS_AND, MODE_ABY, 4 }

    OPCODES[0x0E] = opcode{ INS_ASL, MODE_ABS, 6 }
    OPCODES[0x06] = opcode{ INS_ASL, MODE_ZP0, 5 }
    OPCODES[0x0A] = opcode{ INS_ASL, MODE_ACC, 2 }
    OPCODES[0x16] = opcode{ INS_ASL, MODE_ZPX, 6 }
    OPCODES[0x1E] = opcode{ INS_ASL, MODE_ABX, 6 }

    OPCODES[0x90] = opcode{ INS_BCC, MODE_REL, 2 }

    OPCODES[0xB0] = opcode{ INS_BCS, MODE_REL, 2 }

    OPCODES[0xF0] = opcode{ INS_BEQ, MODE_REL, 2 }

    OPCODES[0x89] = opcode{ INS_BIT, MODE_IMM, 2 }
    OPCODES[0x2C] = opcode{ INS_BIT, MODE_ABS, 4 }
    OPCODES[0x24] = opcode{ INS_BIT, MODE_ZP0, 3 }
    OPCODES[0x34] = opcode{ INS_BIT, MODE_ZPX, 4 }
    OPCODES[0x3C] = opcode{ INS_BIT, MODE_ABX, 4 }

    OPCODES[0x30] = opcode{ INS_BMI, MODE_REL, 2 }

    OPCODES[0xD0] = opcode{ INS_BNE, MODE_REL, 2 }

    OPCODES[0x10] = opcode{ INS_BPL, MODE_REL, 2 }

    OPCODES[0x00] = opcode{ INS_BRK, MODE_IMP, 7 }

    OPCODES[0x50] = opcode{ INS_BVC, MODE_REL, 2 }

    OPCODES[0x70] = opcode{ INS_BVS, MODE_REL, 2 }

    OPCODES[0x18] = opcode{ INS_CLC, MODE_IMP, 2 }

    OPCODES[0xD8] = opcode{ INS_CLD, MODE_IMP, 2 }

    OPCODES[0x58] = opcode{ INS_CLI, MODE_IMP, 2 }

    OPCODES[0xB8] = opcode{ INS_CLV, MODE_IMP, 2 }

    OPCODES[0xC9] = opcode{ INS_CMP, MODE_IMM, 2 }
    OPCODES[0xCD] = opcode{ INS_CMP, MODE_ABS, 4 }
    OPCODES[0xC5] = opcode{ INS_CMP, MODE_ZP0, 3 }
    OPCODES[0xC1] = opcode{ INS_CMP, MODE_INX, 6 }
    OPCODES[0xD1] = opcode{ INS_CMP, MODE_INY, 5 }
    OPCODES[0xD5] = opcode{ INS_CMP, MODE_ZPX, 4 }
    OPCODES[0xDD] = opcode{ INS_CMP, MODE_ABX, 4 }
    OPCODES[0xD9] = opcode{ INS_CMP, MODE_ABY, 4 }

    OPCODES[0xE0] = opcode{ INS_CPX, MODE_IMM, 2 }
    OPCODES[0xEC] = opcode{ INS_CPX, MODE_ABS, 4 }
    OPCODES[0xE4] = opcode{ INS_CPX, MODE_ZP0, 3 }

    OPCODES[0xC0] = opcode{ INS_CPY, MODE_IMM, 2 }
    OPCODES[0xCC] = opcode{ INS_CPY, MODE_ABS, 4 }
    OPCODES[0xC4] = opcode{ INS_CPY, MODE_ZP0, 3 }

    OPCODES[0xCE] = opcode{ INS_DEC, MODE_ABS, 6 }
    OPCODES[0xC6] = opcode{ INS_DEC, MODE_ZP0, 5 }
    OPCODES[0xD6] = opcode{ INS_DEC, MODE_ZPX, 6 }
    OPCODES[0xDE] = opcode{ INS_DEC, MODE_ABX, 6 }

    OPCODES[0xCA] = opcode{ INS_DEX, MODE_IMP, 2 }

    OPCODES[0x88] = opcode{ INS_DEY, MODE_IMP, 2 }

    OPCODES[0x49] = opcode{ INS_EOR, MODE_IMM, 2 }
    OPCODES[0x4D] = opcode{ INS_EOR, MODE_ABS, 4 }
    OPCODES[0x45] = opcode{ INS_EOR, MODE_ZP0, 3 }
    OPCODES[0x41] = opcode{ INS_EOR, MODE_INX, 6 }
    OPCODES[0x51] = opcode{ INS_EOR, MODE_INY, 5 }
    OPCODES[0x55] = opcode{ INS_EOR, MODE_ZPX, 4 }
    OPCODES[0x5D] = opcode{ INS_EOR, MODE_ABX, 4 }
    OPCODES[0x59] = opcode{ INS_EOR, MODE_ABY, 4 }

    OPCODES[0xEE] = opcode{ INS_INC, MODE_ABS, 6 }
    OPCODES[0xE6] = opcode{ INS_INC, MODE_ZP0, 5 }
    OPCODES[0xF6] = opcode{ INS_INC, MODE_ZPX, 6 }
    OPCODES[0xFE] = opcode{ INS_INC, MODE_ABX, 6 }

    OPCODES[0xE8] = opcode{ INS_INX, MODE_IMP, 2 }

    OPCODES[0xC8] = opcode{ INS_INY, MODE_IMP, 2 }

    OPCODES[0x4C] = opcode{ INS_JMP, MODE_ABS, 3 }
    OPCODES[0x6C] = opcode{ INS_JMP, MODE_IND, 6 }

    OPCODES[0x20] = opcode{ INS_JSR, MODE_ABS, 6 }

    OPCODES[0xA9] = opcode{ INS_LDA, MODE_IMM, 2 }
    OPCODES[0xAD] = opcode{ INS_LDA, MODE_ABS, 4 }
    OPCODES[0xA5] = opcode{ INS_LDA, MODE_ZP0, 3 }
    OPCODES[0xA1] = opcode{ INS_LDA, MODE_INX, 6 }
    OPCODES[0xB1] = opcode{ INS_LDA, MODE_INY, 5 }
    OPCODES[0xB5] = opcode{ INS_LDA, MODE_ZPX, 4 }
    OPCODES[0xBD] = opcode{ INS_LDA, MODE_ABX, 4 }
    OPCODES[0xB9] = opcode{ INS_LDA, MODE_ABY, 4 }

    OPCODES[0xA2] = opcode{ INS_LDX, MODE_IMM, 2 }
    OPCODES[0xAE] = opcode{ INS_LDX, MODE_ABS, 4 }
    OPCODES[0xA6] = opcode{ INS_LDX, MODE_ZP0, 3 }
    OPCODES[0xB6] = opcode{ INS_LDX, MODE_ZPY, 4 }
    OPCODES[0xBE] = opcode{ INS_LDX, MODE_ABY, 4 }

    OPCODES[0xA0] = opcode{ INS_LDY, MODE_IMM, 2 }
    OPCODES[0xAC] = opcode{ INS_LDY, MODE_ABS, 4 }
    OPCODES[0xA4] = opcode{ INS_LDY, MODE_ZP0, 3 }
    OPCODES[0xB4] = opcode{ INS_LDY, MODE_ZPX, 4 }
    OPCODES[0xBC] = opcode{ INS_LDY, MODE_ABX, 4 }

    OPCODES[0x4E] = opcode{ INS_LSR, MODE_ABS, 6 }
    OPCODES[0x46] = opcode{ INS_LSR, MODE_ZP0, 5 }
    OPCODES[0x4A] = opcode{ INS_LSR, MODE_ACC, 2 }
    OPCODES[0x56] = opcode{ INS_LSR, MODE_ZPX, 6 }
    OPCODES[0x5E] = opcode{ INS_LSR, MODE_ABX, 6 }

    OPCODES[0xEA] = opcode{ INS_NOP, MODE_IMP, 2 }

    OPCODES[0x09] = opcode{ INS_ORA, MODE_IMM, 2 }
    OPCODES[0x0D] = opcode{ INS_ORA, MODE_ABS, 4 }
    OPCODES[0x05] = opcode{ INS_ORA, MODE_ZP0, 3 }
    OPCODES[0x01] = opcode{ INS_ORA, MODE_INX, 6 }
    OPCODES[0x11] = opcode{ INS_ORA, MODE_INY, 5 }
    OPCODES[0x15] = opcode{ INS_ORA, MODE_ZPX, 4 }
    OPCODES[0x1D] = opcode{ INS_ORA, MODE_ABX, 4 }
    OPCODES[0x19] = opcode{ INS_ORA, MODE_ABY, 4 }

    OPCODES[0x48] = opcode{ INS_PHA, MODE_IMP, 3 }

    OPCODES[0x08] = opcode{ INS_PHP, MODE_IMP, 3 }

    OPCODES[0x68] = opcode{ INS_PLA, MODE_IMP, 4 }

    OPCODES[0x28] = opcode{ INS_PLP, MODE_IMP, 4 }

    OPCODES[0x2E] = opcode{ INS_ROL, MODE_ABS, 6 }
    OPCODES[0x26] = opcode{ INS_ROL, MODE_ZP0, 5 }
    OPCODES[0x2A] = opcode{ INS_ROL, MODE_ACC, 2 }
    OPCODES[0x36] = opcode{ INS_ROL, MODE_ZPX, 6 }
    OPCODES[0x3E] = opcode{ INS_ROL, MODE_ABX, 6 }

    OPCODES[0x6E] = opcode{ INS_ROR, MODE_ABS, 6 }
    OPCODES[0x66] = opcode{ INS_ROR, MODE_ZP0, 5 }
    OPCODES[0x6A] = opcode{ INS_ROR, MODE_ACC, 2 }
    OPCODES[0x76] = opcode{ INS_ROR, MODE_ZPX, 6 }
    OPCODES[0x7E] = opcode{ INS_ROR, MODE_ABX, 6 }

    OPCODES[0x40] = opcode{ INS_RTI, MODE_IMP, 6 }

    OPCODES[0x60] = opcode{ INS_RTS, MODE_IMP, 6 }

    OPCODES[0xE9] = opcode{ INS_SBC, MODE_IMM, 2 }
    OPCODES[0xED] = opcode{ INS_SBC, MODE_ABS, 4 }
    OPCODES[0xE5] = opcode{ INS_SBC, MODE_ZP0, 3 }
    OPCODES[0xE1] = opcode{ INS_SBC, MODE_INX, 6 }
    OPCODES[0xF1] = opcode{ INS_SBC, MODE_INY, 5 }
    OPCODES[0xF5] = opcode{ INS_SBC, MODE_ZPX, 4 }
    OPCODES[0xFD] = opcode{ INS_SBC, MODE_ABX, 4 }
    OPCODES[0xF9] = opcode{ INS_SBC, MODE_ABY, 4 }

    OPCODES[0x38] = opcode{ INS_SEC, MODE_IMP, 2 }

    OPCODES[0xF8] = opcode{ INS_SED, MODE_IMP, 2 }

    OPCODES[0x78] = opcode{ INS_SEI, MODE_IMP, 2 }

    OPCODES[0x8D] = opcode{ INS_STA, MODE_ABS, 4 }
    OPCODES[0x85] = opcode{ INS_STA, MODE_ZP0, 3 }
    OPCODES[0x81] = opcode{ INS_STA, MODE_INX, 6 }
    OPCODES[0x91] = opcode{ INS_STA, MODE_INY, 6 }
    OPCODES[0x95] = opcode{ INS_STA, MODE_ZPX, 4 }
    OPCODES[0x9D] = opcode{ INS_STA, MODE_ABX, 5 }
    OPCODES[0x99] = opcode{ INS_STA, MODE_ABY, 5 }

    OPCODES[0x8E] = opcode{ INS_STX, MODE_ABS, 4 }
    OPCODES[0x86] = opcode{ INS_STX, MODE_ZP0, 3 }
    OPCODES[0x96] = opcode{ INS_STX, MODE_ZPY, 4 }

    OPCODES[0x8C] = opcode{ INS_STY, MODE_ABS, 4 }
    OPCODES[0x84] = opcode{ INS_STY, MODE_ZP0, 3 }
    OPCODES[0x94] = opcode{ INS_STY, MODE_ZPX, 4 }

    OPCODES[0xAA] = opcode{ INS_TAX, MODE_IMP, 2 }

    OPCODES[0xA8] = opcode{ INS_TAY, MODE_IMP, 2 }

    OPCODES[0xBA] = opcode{ INS_TSX, MODE_IMP, 2 }

    OPCODES[0x8A] = opcode{ INS_TXA, MODE_IMP, 2 }

    OPCODES[0x9A] = opcode{ INS_TXS, MODE_IMP, 2 }

    OPCODES[0x98] = opcode{ INS_TYA, MODE_IMP, 2 }
}
