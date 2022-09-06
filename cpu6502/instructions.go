package cpu6502

type Instruction string

const (
	INS_ADC Instruction = "ADC"
	INS_AND Instruction = "AND"
	INS_ASL Instruction = "ASL"
	INS_BCC Instruction = "BCC"
	INS_BCS Instruction = "BCS"
	INS_BEQ Instruction = "BEQ"
	INS_BIT Instruction = "BIT"
	INS_BMI Instruction = "BMI"
	INS_BNE Instruction = "BNE"
	INS_BPL Instruction = "BPL"
	INS_BRK Instruction = "BRK"
	INS_BVC Instruction = "BVC"
	INS_BVS Instruction = "BVS"
	INS_CLC Instruction = "CLC"
	INS_CLD Instruction = "CLD"
	INS_CLI Instruction = "CLI"
	INS_CLV Instruction = "CLV"
	INS_CMP Instruction = "CMP"
	INS_CPX Instruction = "CPX"
	INS_CPY Instruction = "CPY"
	INS_DEC Instruction = "DEC"
	INS_DEX Instruction = "DEX"
	INS_DEY Instruction = "DEY"
	INS_EOR Instruction = "EOR"
	INS_INC Instruction = "INC"
	INS_INX Instruction = "INX"
	INS_INY Instruction = "INY"
	INS_JMP Instruction = "JMP"
	INS_JSR Instruction = "JSR"
	INS_LDA Instruction = "LDA"
	INS_LDX Instruction = "LDX"
	INS_LDY Instruction = "LDY"
	INS_LSR Instruction = "LSR"
	INS_NOP Instruction = "NOP"
	INS_ORA Instruction = "ORA"
	INS_PHA Instruction = "PHA"
	INS_PHP Instruction = "PHP"
	INS_PLA Instruction = "PLA"
	INS_PLP Instruction = "PLP"
	INS_ROL Instruction = "ROL"
	INS_ROR Instruction = "ROR"
	INS_RTI Instruction = "RTI"
	INS_RTS Instruction = "RTS"
	INS_SBC Instruction = "SBC"
	INS_SEC Instruction = "SEC"
	INS_SED Instruction = "SED"
	INS_SEI Instruction = "SEI"
	INS_STA Instruction = "STA"
	INS_STX Instruction = "STX"
	INS_STY Instruction = "STY"
	INS_TAX Instruction = "TAX"
	INS_TAY Instruction = "TAY"
	INS_TSX Instruction = "TSX"
	INS_TXA Instruction = "TXA"
	INS_TXS Instruction = "TXS"
	INS_TYA Instruction = "TYA"
)

func attachInstructions(cpu *CPU) {
	cpu.instructions = make(Instructions)

	cpu.instructions[INS_ADC] = cpu.adc
	cpu.instructions[INS_AND] = cpu.and
	cpu.instructions[INS_ASL] = cpu.asl
	cpu.instructions[INS_BCC] = cpu.bcc
	cpu.instructions[INS_BCS] = cpu.bcs
	cpu.instructions[INS_BEQ] = cpu.beq
	cpu.instructions[INS_BIT] = cpu.bit
	cpu.instructions[INS_BMI] = cpu.bmi
	cpu.instructions[INS_BNE] = cpu.bne
	cpu.instructions[INS_BPL] = cpu.bpl
	cpu.instructions[INS_BRK] = cpu.brk
	cpu.instructions[INS_BVC] = cpu.bvc
	cpu.instructions[INS_BVS] = cpu.bvs
	cpu.instructions[INS_CLC] = cpu.clc
	cpu.instructions[INS_CLD] = cpu.cld
	cpu.instructions[INS_CLI] = cpu.cli
	cpu.instructions[INS_CLV] = cpu.clv
	cpu.instructions[INS_CMP] = cpu.cmp
	cpu.instructions[INS_CPX] = cpu.cpx
	cpu.instructions[INS_CPY] = cpu.cpy
	cpu.instructions[INS_DEC] = cpu.dec
	cpu.instructions[INS_DEX] = cpu.dex
	cpu.instructions[INS_DEY] = cpu.dey
	cpu.instructions[INS_EOR] = cpu.eor
	cpu.instructions[INS_INC] = cpu.inc
	cpu.instructions[INS_INX] = cpu.incx
	cpu.instructions[INS_INY] = cpu.incy
	cpu.instructions[INS_JMP] = cpu.jmp
	cpu.instructions[INS_JSR] = cpu.jsr
	cpu.instructions[INS_LDA] = cpu.lda
	cpu.instructions[INS_LDX] = cpu.ldx
	cpu.instructions[INS_LDY] = cpu.ldy
	cpu.instructions[INS_LSR] = cpu.lsr
	cpu.instructions[INS_NOP] = cpu.nop
	cpu.instructions[INS_ORA] = cpu.ora
	cpu.instructions[INS_PHA] = cpu.pha
	cpu.instructions[INS_PHP] = cpu.php
	cpu.instructions[INS_PLA] = cpu.pla
	cpu.instructions[INS_PLP] = cpu.plp
	cpu.instructions[INS_ROL] = cpu.rol
	cpu.instructions[INS_ROR] = cpu.ror
	cpu.instructions[INS_RTI] = cpu.rti
	cpu.instructions[INS_RTS] = cpu.rts
	cpu.instructions[INS_SBC] = cpu.sbc
	cpu.instructions[INS_SEC] = cpu.sec
	cpu.instructions[INS_SED] = cpu.sed
	cpu.instructions[INS_SEI] = cpu.sei
	cpu.instructions[INS_STA] = cpu.sta
	cpu.instructions[INS_STX] = cpu.stx
	cpu.instructions[INS_STY] = cpu.sty
	cpu.instructions[INS_TAX] = cpu.tax
	cpu.instructions[INS_TAY] = cpu.tay
	cpu.instructions[INS_TSX] = cpu.tsx
	cpu.instructions[INS_TXA] = cpu.txa
	cpu.instructions[INS_TXS] = cpu.txs
	cpu.instructions[INS_TYA] = cpu.tya
}

// Loads data from the address depending on the address mode
func (cpu *CPU) loadData(mode AddressingMode) (byte, uint16) {
	address := cpu.addressingModes[mode]()

	if mode == MODE_ACC {
		return cpu.A, address
	}

	if mode == MODE_IMP {
		return 0, address
	}

	if mode == MODE_REL {
		offset := cpu.read(address)

		// since the operand could be a negative number we need to verify if the
		// most significant bit on the left is 1 and then convert it to uint16 properly
		// so it can be added to the Program Counter correctly
		if offset&0x80 > 0 {
			return 0, 0xFF00 | uint16(offset)
		}
		return 0, uint16(offset)
	}

	return cpu.read(address), address
}

// write data back to the address depending on the AddressMode
func (cpu *CPU) writeData(data byte, address uint16, mode AddressingMode) {
	if mode == MODE_ACC {
		cpu.A = data
		return
	}

	cpu.write(address, data)
}

// Addition
// Add memory to accumulator with carry bit
// It sets the overflow flag when there is a two's complement overflow
// considering teh offset -128 to +127 as it works with 8 bits
// It bases the overflow formula on the accordingly truth table for the most significant bit of the operation
//
//	        A | M | R | V  | A^R | M^R
//	        0   0   0   0     0     0
//	        0   0   1   1     1     1
//	        0   1   0   0     0     1
//	        0   1   1   0     1     0   So Overflow formula = (A^R) & (M^R)
//	        1   1   0   1     1     1
//	        1   1   1   0     0     0
//	        1   0   1   0     0     1
//	        1   0   0   0     1     0
//
//	Where A = Accumulator, M = Read memory, R = Result, V = Should overflow
func (cpu *CPU) adc(mode AddressingMode) {
	data, _ := cpu.loadData(mode)

	result := uint16(cpu.A) + uint16(data) + uint16(cpu.GetFlag(FLAG_C))

	cpu.SetFlag(FLAG_Z, (result&0x00FF) == 0x0000)
	cpu.SetFlag(FLAG_N, result&0x0080 > 0)
	cpu.SetFlag(FLAG_C, result&0xFF00 > 0)

	overflow := ((cpu.A ^ byte(result&0x00FF)) & (data ^ byte(result&0x00FF))) & 0x80

	cpu.SetFlag(FLAG_V, overflow > 0)

	cpu.A = byte(result & 0x00FF)
}

// Subtract memory from accumulator with borrow
func (cpu *CPU) sbc(mode AddressingMode) {
	data, _ := cpu.loadData(mode)

	// uses two complement to get the inverse signal of the memory value
	data = (^data) + 1

	result := uint16(cpu.A) + uint16(data) + uint16(cpu.GetFlag(FLAG_C))

	cpu.SetFlag(FLAG_Z, (result&0x00FF) == 0x0000)
	cpu.SetFlag(FLAG_N, result&0x0080 > 0)
	cpu.SetFlag(FLAG_C, result&0xFF00 > 0)

	overflow := ((cpu.A ^ byte(result&0x00FF)) & (data ^ byte(result&0x00FF))) & 0x80

	cpu.SetFlag(FLAG_V, overflow > 0)

	cpu.A = byte(result & 0x00FF)
}

// And memory with accumulator
func (cpu *CPU) and(mode AddressingMode) {
	data, _ := cpu.loadData(mode)
	cpu.A = cpu.A & data

	cpu.SetFlag(FLAG_Z, cpu.A == 0x00)
	cpu.SetFlag(FLAG_N, cpu.A&0x80 > 0)
}

// Shift left one bit ( Memory or accumulator )
func (cpu *CPU) asl(mode AddressingMode) {
	data, address := cpu.loadData(mode)

	shifted := uint16(data) << 1

	cpu.SetFlag(FLAG_Z, (shifted&0x00FF) == 0x0000)
	cpu.SetFlag(FLAG_N, shifted&0x0080 > 0)
	cpu.SetFlag(FLAG_C, shifted&0xFF00 > 0)

	cpu.writeData(byte(shifted&0x00FF), address, mode)
}

// Branch on carry clear
func (cpu *CPU) bcc(mode AddressingMode) {
	if cpu.GetFlag(FLAG_C) != 0x00 {
		return
	}

	_, offset := cpu.loadData(mode)
	cpu.PC += offset
}

// Branch on carry set
func (cpu *CPU) bcs(mode AddressingMode) {
	if cpu.GetFlag(FLAG_C) == 0x00 {
		return
	}

	_, offset := cpu.loadData(mode)
	cpu.PC += offset

}

// Branch on result zero (when zero flag set)
func (cpu *CPU) beq(mode AddressingMode) {
	if cpu.GetFlag(FLAG_Z) == 0x00 {
		return
	}

	_, offset := cpu.loadData(mode)
	cpu.PC += offset
}

// Test bits in memory with Accumulator
func (cpu *CPU) bit(mode AddressingMode) {
	data, _ := cpu.loadData(mode)
	result := cpu.A & data

	cpu.SetFlag(FLAG_Z, result == 0x00)
	cpu.SetFlag(FLAG_N, data&0x40 > 0) // Memory bit 7
	cpu.SetFlag(FLAG_V, data&0x20 > 0) // Memory bit 6
}

// Branch on result minus (when negative flag set)
func (cpu *CPU) bmi(mode AddressingMode) {
	if cpu.GetFlag(FLAG_N) == 0x00 {
		return
	}

	_, offset := cpu.loadData(mode)
	cpu.PC += offset
}

// Branch on result not zero (when zero flag not set)
func (cpu *CPU) bne(mode AddressingMode) {
	if cpu.GetFlag(FLAG_Z) != 0x00 {
		return
	}

	_, offset := cpu.loadData(mode)
	cpu.PC += offset
}

// Branch on result plus (when negative flag not set)
func (cpu *CPU) bpl(mode AddressingMode) {
	if cpu.GetFlag(FLAG_N) != 0x00 {
		return
	}

	_, offset := cpu.loadData(mode)
	cpu.PC += offset
}

// Break
// Stores the Program Counter and the Status on the stack
// Loads the PC from low = 0xFFFE, high = 0xFFFF
func (cpu *CPU) brk(mode AddressingMode) {
	cpu.SetFlag(FLAG_I, true)

	pcl := byte(cpu.PC & 0x00FF)
	pch := byte((cpu.PC >> 8) & 0x00FF)

	cpu.pushOnStack(pch)
	cpu.pushOnStack(pcl)

	cpu.SetFlag(FLAG_B, true)
	cpu.pushOnStack(cpu.Status)
	cpu.SetFlag(FLAG_B, false)

	low := uint16(cpu.read(0xFFFE))
	high := uint16(cpu.read(0xFFFF))

	cpu.PC = (high << 8) | low

}

// Branch on overflow clear (when overflow flag is not set)
func (cpu *CPU) bvc(mode AddressingMode) {
	if cpu.GetFlag(FLAG_V) != 0x00 {
		return
	}

	_, offset := cpu.loadData(mode)
	cpu.PC += offset
}

// Branch on overflow set (when overflow flag set)
func (cpu *CPU) bvs(mode AddressingMode) {
	if cpu.GetFlag(FLAG_V) == 0x00 {
		return
	}

	_, offset := cpu.loadData(mode)
	cpu.PC += offset
}

// Clears carry flag
func (cpu *CPU) clc(mode AddressingMode) {
	cpu.SetFlag(FLAG_C, false)
}

// Clears decimal mode
func (cpu *CPU) cld(mode AddressingMode) {
	cpu.SetFlag(FLAG_D, false)
}

// Clears interrupt disabled bit
func (cpu *CPU) cli(mode AddressingMode) {
	cpu.SetFlag(FLAG_I, false)
}

// Clears overflow flag
func (cpu *CPU) clv(mode AddressingMode) {
	cpu.SetFlag(FLAG_C, false)
}

// Compare memory with accumulator
func (cpu *CPU) cmp(mode AddressingMode) {
	data, _ := cpu.loadData(mode)
	result := cpu.A - data

	cpu.SetFlag(FLAG_Z, result == 0x00)
	cpu.SetFlag(FLAG_N, result&0x80 > 0)
	cpu.SetFlag(FLAG_C, cpu.A >= data)
}

// Compare memory with X register
func (cpu *CPU) cpx(mode AddressingMode) {
	data, _ := cpu.loadData(mode)
	result := cpu.X - data

	cpu.SetFlag(FLAG_Z, result == 0x00)
	cpu.SetFlag(FLAG_N, result&0x80 > 0)
	cpu.SetFlag(FLAG_C, cpu.X >= data)
}

// Compare memory with Y register
func (cpu *CPU) cpy(mode AddressingMode) {
	data, _ := cpu.loadData(mode)
	result := cpu.Y - data

	cpu.SetFlag(FLAG_Z, result == 0x00)
	cpu.SetFlag(FLAG_N, result&0x80 > 0)
	cpu.SetFlag(FLAG_C, cpu.Y >= data)
}

// Decrement memory by 1
func (cpu *CPU) dec(mode AddressingMode) {
	data, address := cpu.loadData(mode)
	result := data - 1

	cpu.SetFlag(FLAG_Z, result == 0x00)
	cpu.SetFlag(FLAG_N, result&0x80 > 0)

	cpu.writeData(result, address, mode)
}

// Decrement X Register by 1
func (cpu *CPU) dex(mode AddressingMode) {
	cpu.X -= 1

	cpu.SetFlag(FLAG_Z, cpu.X == 0x00)
	cpu.SetFlag(FLAG_N, cpu.X&0x80 > 0)
}

// Decrement Y Register by 1
func (cpu *CPU) dey(mode AddressingMode) {
	cpu.Y -= 1

	cpu.SetFlag(FLAG_Z, cpu.Y == 0x00)
	cpu.SetFlag(FLAG_N, cpu.Y&0x80 > 0)
}

// Exclusive or with memory and accumulator
func (cpu *CPU) eor(mode AddressingMode) {
	data, _ := cpu.loadData(mode)

	cpu.A = cpu.A ^ data

	cpu.SetFlag(FLAG_Z, cpu.A == 0x00)
	cpu.SetFlag(FLAG_N, cpu.A&0x80 > 0)
}

// Increment memory by 1
func (cpu *CPU) inc(mode AddressingMode) {
	data, address := cpu.loadData(mode)
	result := data + 1

	cpu.SetFlag(FLAG_Z, result == 0x00)
	cpu.SetFlag(FLAG_N, result&0x80 > 0)

	cpu.writeData(result, address, mode)
}

// Increment X Register by 1
func (cpu *CPU) incx(mode AddressingMode) {
	cpu.X += 1

	cpu.SetFlag(FLAG_Z, cpu.X == 0x00)
	cpu.SetFlag(FLAG_N, cpu.X&0x80 > 0)
}

// Increment Y Register by 1
func (cpu *CPU) incy(mode AddressingMode) {
	cpu.Y += 1

	cpu.SetFlag(FLAG_Z, cpu.Y == 0x00)
	cpu.SetFlag(FLAG_N, cpu.Y&0x80 > 0)
}

// Jump to new location
func (cpu *CPU) jmp(mode AddressingMode) {
	_, address := cpu.loadData(mode)

	cpu.PC = address
}

// Jump to subroutine
func (cpu *CPU) jsr(mode AddressingMode) {
	_, address := cpu.loadData(mode)

	// Store the last byte address from the current
	// operation to the stack
	cpu.PC--

	pcl := byte(cpu.PC & 0x00FF)
	pch := byte((cpu.PC >> 8) & 0x00FF)

	cpu.pushOnStack(pch)
	cpu.pushOnStack(pcl)

	cpu.PC = address
}

// Load accumulator with memory
func (cpu *CPU) lda(mode AddressingMode) {
	data, _ := cpu.loadData(mode)
	cpu.A = data

	cpu.SetFlag(FLAG_Z, cpu.A == 0x00)
	cpu.SetFlag(FLAG_N, cpu.A&0x80 > 0)
}

// Load index X with memory
func (cpu *CPU) ldx(mode AddressingMode) {
	data, _ := cpu.loadData(mode)
	cpu.X = data

	cpu.SetFlag(FLAG_Z, cpu.X == 0x00)
	cpu.SetFlag(FLAG_N, cpu.X&0x80 > 0)
}

// Load index Y with memory
func (cpu *CPU) ldy(mode AddressingMode) {
	data, _ := cpu.loadData(mode)
	cpu.Y = data

	cpu.SetFlag(FLAG_Z, cpu.Y == 0x00)
	cpu.SetFlag(FLAG_N, cpu.Y&0x80 > 0)
}

// Shift one bit right (memory or accumulator)
func (cpu *CPU) lsr(mode AddressingMode) {
	data, address := cpu.loadData(mode)

	result := data >> 1

	cpu.SetFlag(FLAG_Z, result == 0x00)
	cpu.SetFlag(FLAG_N, result&0x80 > 0)
	cpu.SetFlag(FLAG_C, data&0x01 > 0)

	cpu.writeData(result, address, mode)
}

// No operation
func (cpu *CPU) nop(mode AddressingMode) {

}

// OR memory with accumulator
func (cpu *CPU) ora(mode AddressingMode) {
	data, _ := cpu.loadData(mode)

	cpu.A = cpu.A | data

	cpu.SetFlag(FLAG_Z, cpu.A == 0x00)
	cpu.SetFlag(FLAG_N, cpu.A&0x80 > 0)
}

// Push accumulator on Stack
func (cpu *CPU) pha(mode AddressingMode) {
	cpu.pushOnStack(cpu.A)
}

// Push processor status on stack
func (cpu *CPU) php(mode AddressingMode) {
	cpu.pushOnStack(cpu.Status)
}

// Pull accumulator from Stack
func (cpu *CPU) pla(mode AddressingMode) {
	cpu.A = cpu.pullFromStack()

	cpu.SetFlag(FLAG_Z, cpu.A == 0x00)
	cpu.SetFlag(FLAG_N, cpu.A&0x80 > 0)
}

// Pull processor status from Stack
func (cpu *CPU) plp(mode AddressingMode) {
	cpu.Status = cpu.pullFromStack()
}

// Rotate one bit left (memory or accumulator)
func (cpu *CPU) rol(mode AddressingMode) {
	data, address := cpu.loadData(mode)

	cpu.SetFlag(FLAG_C, data&0x80 > 0)

	result := (data << 1) | cpu.GetFlag(FLAG_C)

	cpu.SetFlag(FLAG_Z, result == 0x00)
	cpu.SetFlag(FLAG_N, result&0x80 > 0)

	cpu.writeData(result, address, mode)
}

// Rotate one bit right (memory or accumulator)
func (cpu *CPU) ror(mode AddressingMode) {
	data, address := cpu.loadData(mode)

	cpu.SetFlag(FLAG_C, data&0x01 > 0)

	result := (data >> 1) | (cpu.GetFlag(FLAG_C) << 7)

	cpu.SetFlag(FLAG_Z, result == 0x00)
	cpu.SetFlag(FLAG_N, result&0x80 > 0)

	cpu.writeData(result, address, mode)
}

// Return from interruption
func (cpu *CPU) rti(mode AddressingMode) {
	cpu.Status = cpu.pullFromStack()

	low := uint16(cpu.pullFromStack())
	high := uint16(cpu.pullFromStack())

	cpu.PC = (high << 8) | low
}

// Return from subroutine
func (cpu *CPU) rts(mode AddressingMode) {
	low := uint16(cpu.pullFromStack())
	high := uint16(cpu.pullFromStack())

	cpu.PC = (high << 8) | low
	cpu.PC++
}

// Set carry flag
func (cpu *CPU) sec(mode AddressingMode) {
	cpu.SetFlag(FLAG_C, true)
}

// Set decimal mode
func (cpu *CPU) sed(mode AddressingMode) {
	cpu.SetFlag(FLAG_D, true)
}

// Set disable interrupts
func (cpu *CPU) sei(mode AddressingMode) {
	cpu.SetFlag(FLAG_I, true)
}

// Store accumulator in memory
func (cpu *CPU) sta(mode AddressingMode) {
	_, address := cpu.loadData(mode)
	cpu.write(address, cpu.A)
}

// Store X register in memory
func (cpu *CPU) stx(mode AddressingMode) {
	_, address := cpu.loadData(mode)
	cpu.write(address, cpu.X)
}

// Store Y register in memory
func (cpu *CPU) sty(mode AddressingMode) {
	_, address := cpu.loadData(mode)
	cpu.write(address, cpu.Y)
}

// Transfer accumulator to X register
func (cpu *CPU) tax(mode AddressingMode) {
	cpu.X = cpu.A

	cpu.SetFlag(FLAG_Z, cpu.X == 0x00)
	cpu.SetFlag(FLAG_N, cpu.X&0x80 > 0)
}

// Transfer accumulator to Y register
func (cpu *CPU) tay(mode AddressingMode) {
	cpu.Y = cpu.A

	cpu.SetFlag(FLAG_Z, cpu.Y == 0x00)
	cpu.SetFlag(FLAG_N, cpu.Y&0x80 > 0)
}

// Transfer stack pointer to index X
func (cpu *CPU) tsx(mode AddressingMode) {
	cpu.X = cpu.S

	cpu.SetFlag(FLAG_Z, cpu.X == 0x00)
	cpu.SetFlag(FLAG_N, cpu.X&0x80 > 0)
}

// Transfer X register to accumulator
func (cpu *CPU) txa(mode AddressingMode) {
	cpu.A = cpu.X

	cpu.SetFlag(FLAG_Z, cpu.A == 0x00)
	cpu.SetFlag(FLAG_N, cpu.A&0x80 > 0)
}

// Transfer X register to stack pointer
func (cpu *CPU) txs(mode AddressingMode) {
	cpu.S = cpu.X
}

// Transfer Y register to accumulator
func (cpu *CPU) tya(mode AddressingMode) {
	cpu.A = cpu.Y

	cpu.SetFlag(FLAG_Z, cpu.Y == 0x00)
	cpu.SetFlag(FLAG_N, cpu.Y&0x80 > 0)
}
