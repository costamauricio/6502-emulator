package cpu6502

type AddressingMode string

const (
    MODE_ACC AddressingMode = "ACC" // Accumulator
    MODE_IMM AddressingMode = "IMM" // Immediate
    MODE_ABS AddressingMode = "ABS" // Absolute
    MODE_ZP0 AddressingMode = "ZP" // ZeroPage
    MODE_ZPX AddressingMode = "ZP, X" // ZeroPage, X
    MODE_ZPY AddressingMode = "ZP, Y" // ZeroPage, Y
    MODE_ABX AddressingMode = "ABS, X" // Absolute, X
    MODE_ABY AddressingMode = "ABS, Y" // Absolute, Y
    MODE_IMP AddressingMode = "IMP" // Implied
    MODE_REL AddressingMode = "REL" // Relative
    MODE_IND AddressingMode = "IND" // Indirect
    MODE_INX AddressingMode = "IND, X" // Indirect, X
    MODE_INY AddressingMode = "IND, Y" // Indirect, Y
)

func attachAddressModes(cpu *CPU) {
    cpu.addressingModes = make(AddressingModes)

    cpu.addressingModes[MODE_ACC] = cpu.acc
    cpu.addressingModes[MODE_IMM] = cpu.imm
    cpu.addressingModes[MODE_ABS] = cpu.abs
    cpu.addressingModes[MODE_ZP0] = cpu.zp0
    cpu.addressingModes[MODE_ZPX] = cpu.zpx
    cpu.addressingModes[MODE_ZPY] = cpu.zpy
    cpu.addressingModes[MODE_ABX] = cpu.abx
    cpu.addressingModes[MODE_ABY] = cpu.aby
    cpu.addressingModes[MODE_IMP] = cpu.imp
    cpu.addressingModes[MODE_REL] = cpu.rel
    cpu.addressingModes[MODE_IND] = cpu.ind
    cpu.addressingModes[MODE_INX] = cpu.inx
    cpu.addressingModes[MODE_INY] = cpu.iny
}

func (cpu *CPU) acc() uint16 {
    return 0
}

// Immediate addressing mode
// The second byte of the instruction is the needed address
func (cpu *CPU) imm() uint16 {
    address := cpu.PC
    cpu.PC++

    return address
}

// Absolute addressing mode
// The second byte of the instruction contains the low order bits
// and the third byte contains the high order bits of the effective address
func (cpu *CPU) abs() uint16 {
    low := uint16(cpu.read(cpu.PC))
    cpu.PC++

    high := uint16(cpu.read(cpu.PC))
    cpu.PC++

    return (high<<8) | low
}

// Zero page addressing mode
// The second byte of the instruction contains the low order bits
// and is assumed a zero high order bits (Zero Page)
func (cpu *CPU) zp0() uint16 {
    address := cpu.read(cpu.PC)
    cpu.PC++

    return uint16(address)
}

// Zero page, X addressing mode (Indexed Zero Page)
// The second byte of the instruction is added to the contents of the X register
// to compose the low order bits
// it is assumed a zero high order bits (Zero Page)
func (cpu *CPU) zpx() uint16 {
    address := cpu.read(cpu.PC) + cpu.X
    cpu.PC++

    return uint16(address)
}

// Zero page, Y addressing mode (Indexed Zero Page)
// The second byte of the instruction is added to the contents of the Y register
// to compose the low order bits
// it is assumed a zero high order bits (Zero Page)
func (cpu *CPU) zpy() uint16 {
    address := cpu.read(cpu.PC) + cpu.Y
    cpu.PC++

    return uint16(address)
}

// Absolute, X addressing mode (Indexed Absolute)
// The second byte of the instruction contains the low order bits
// and the third byte contains the high order bits of the effective address
// Both are indexed by the X register
func (cpu *CPU) abx() uint16 {
    low := uint16(cpu.read(cpu.PC))
    cpu.PC++

    high := uint16(cpu.read(cpu.PC))
    cpu.PC++

    return ((high<<8) | low) + uint16(cpu.X)
}

// Absolute, Y addressing mode (Indexed Absolute)
// The second byte of the instruction contains the low order bits
// and the third byte contains the high order bits of the effective address
// Both are indexed by the Y register
func (cpu *CPU) aby() uint16 {
    low := uint16(cpu.read(cpu.PC))
    cpu.PC++

    high := uint16(cpu.read(cpu.PC))
    cpu.PC++

    return ((high<<8) | low) + uint16(cpu.Y)
}

// Implied addressing mode
func (cpu *CPU) imp() uint16 {
    return 0;
}

// Relative addressing mode
// Used for branch instructions
// The second byte of the instruction contains is an offset that will be added to the Program Counter
// The range of the offset is -128 to +127 bytes
func (cpu *CPU) rel() uint16 {
    address := cpu.PC
    cpu.PC++

    return uint16(address)
}

// Absolute indirect adressing mode
// The second byte of the instruction contains a low order bits of a memory address
// The third byte contains the high order bits
// The contents of the memory address contains the low order bits of the effective address
// The next memory location contains the high order bits
func (cpu *CPU) ind() uint16 {
    low := uint16(cpu.read(cpu.PC))
    cpu.PC++

    high := uint16(cpu.read(cpu.PC))
    cpu.PC++

    pointer := (high<<8) | low

    low = uint16(cpu.read(pointer))
    high = uint16(cpu.read(pointer+1))

    return (high<<8) | low
}

// Indexed indirect addressing mode (Indirect X)
// The second byte of the instruction is added to the contents of X discarding the carry bit
// The result points to a memory location on page 0 (0x00) that contains the low order bits of the address
// The next memory address contains the high order bits
func (cpu *CPU) inx() uint16 {
    location := uint16(cpu.read(cpu.PC) + cpu.X) & 0x00FF
    cpu.PC++

    low := uint16(cpu.read(location))
    high := uint16(cpu.read(location+1))

    return (high<<8) | low
}

// Indirect indexed addressing mode (Indirect Y)
// The second byte of the instruction is the low order bits of a address in page zero
// The contents of that address is the low order bits and the next will be the high order bits
// Then the contents of the Y register are added to the formed address to get the effective address
func (cpu *CPU) iny() uint16 {
    location := uint16(cpu.read(cpu.PC))
    cpu.PC++

    low := uint16(cpu.read(location))
    high := uint16(cpu.read(location+1)) & 0x00FF

    address := (high<<8) | low
    address += uint16(cpu.Y)

    return address
}

