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

    OPCODES[0xA9] = opcode{ INS_LDA, MODE_IMM, 2 }
    OPCODES[0xAD] = opcode{ INS_LDA, MODE_ABS, 4 }
    OPCODES[0xA5] = opcode{ INS_LDA, MODE_ZP0, 3 }
    OPCODES[0xA1] = opcode{ INS_LDA, MODE_INX, 6 }
    OPCODES[0xB1] = opcode{ INS_LDA, MODE_INY, 5 }
    OPCODES[0xB5] = opcode{ INS_LDA, MODE_ZPX, 4 }
    OPCODES[0xBD] = opcode{ INS_LDA, MODE_ABX, 4 }
    OPCODES[0xB9] = opcode{ INS_LDA, MODE_ABY, 4 }

    OPCODES[0x00] = opcode{ INS_BRK, MODE_IMP, 7 }
}
