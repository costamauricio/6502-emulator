package cpu6502

import (
	"fmt"
)

func (cpu *CPU) DisassembleInstructions(startAt uint16, endAt uint16) (map[uint16]string, []uint16) {
	var order []uint16
	instructions := make(map[uint16]string)

	for address := uint(startAt); address <= uint(endAt); address++ {
		opcode := cpu.read(uint16(address))

		operation, found := OPCODES[opcode]

		if !found {
			instructions[uint16(address)] = "UNKNOWN"
			continue
		}

		instructionLocation := address
		var parsed string

		parsed = fmt.Sprintf("$%04X:  ", address) + string(operation.instruction) + " "

		cpu.PC = uint16(address + 1)
		content, finalAddress := cpu.loadData(operation.addressMode)

		switch operation.addressMode {
		case MODE_ACC:
			parsed += "A"
		case MODE_IMM:
			parsed += "#" + fmt.Sprintf("%02X", content)
			address++
		case MODE_ABS, MODE_ABX, MODE_ABY, MODE_ZP0, MODE_ZPX, MODE_ZPY:
			parsed += fmt.Sprintf("$%04X  (%s)", finalAddress, string(operation.addressMode))
			address = uint(cpu.PC - 1)
		case MODE_IMP:
			parsed += " (" + string(operation.addressMode) + ")"
		case MODE_REL:
			parsed += fmt.Sprintf("$%02X  [$%04X]  (%s)", byte(finalAddress&0x00FF), uint16(address)+finalAddress, string(operation.addressMode))
			address = uint(cpu.PC - 1)
		}

		instructions[uint16(instructionLocation)] = parsed
		order = append(order, uint16(instructionLocation))
	}

	return instructions, order
}
