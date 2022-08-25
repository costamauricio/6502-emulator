package cpu6502

type AddressingMode func(*CPU) uint16

// Immediate addressing mode
// The second byte of the instruction is the needed address
func imm(cpu *CPU) uint16 {
    address := cpu.PC
    cpu.PC++

    return address
}

// Absolute addressing mode
// The second byte of the instruction contains the low order bits
// and the third byte contains the high order bits of the effective address
func abs(cpu *CPU) uint16 {
    low := uint16(cpu.read(cpu.PC))
    cpu.PC++

    high := uint16(cpu.read(cpu.PC))
    cpu.PC++

    return (high<<8) | low
}

// Zero page addressing mode
// The second byte of the instruction contains the low order bits
// and is assumed a zero high order bits (Zero Page)
func zp0(cpu *CPU) uint16 {
    address := cpu.read(cpu.PC)
    cpu.PC++

    return uint16(address)
}

// Zero page, X addressing mode (Indexed Zero Page)
// The second byte of the instruction is added to the contents of the X register
// to compose the low order bits
// it is assumed a zero high order bits (Zero Page)
func zpx(cpu *CPU) uint16 {
    address := cpu.read(cpu.PC) + cpu.X
    cpu.PC++

    return uint16(address)
}

// Zero page, Y addressing mode (Indexed Zero Page)
// The second byte of the instruction is added to the contents of the Y register
// to compose the low order bits
// it is assumed a zero high order bits (Zero Page)
func zpy(cpu *CPU) uint16 {
    address := cpu.read(cpu.PC) + cpu.Y
    cpu.PC++

    return uint16(address)
}

// Absolute, X addressing mode (Indexed Absolute)
// The second byte of the instruction contains the low order bits
// and the third byte contains the high order bits of the effective address
// Both are indexed by the X register
func abx(cpu *CPU) uint16 {
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
func aby(cpu *CPU) uint16 {
    low := uint16(cpu.read(cpu.PC))
    cpu.PC++

    high := uint16(cpu.read(cpu.PC))
    cpu.PC++

    return ((high<<8) | low) + uint16(cpu.Y)
}

// Implied addressing mode
func imp(cpu *CPU) uint16 {
    return 0;
}

// Relative addressing mode
// Used for branch instructions
// The second byte of the instruction is an offset that will be added to the Program Counter
// The range of the offset is -127 to +127 bytes
func rel(cpu *CPU) uint16 {
    address := cpu.PC
    cpu.PC++

    return address
    //offset := cpu.read(cpu.PC)
    //cpu.PC++

    //// since the operand could be a negative number we need to verify if the
    //// most significant bit on the left is 1 and then convert it to uint16 properly
    //if offset & 0x80 > 0 {
        //return uint16(offset) | 0xFF00
    //}

    //return uint16(offset)
}

// Absolute indirect adressing mode
// The second byte of the instruction contains a low order bits of a memory address
// The third byte contains the high order bits
// The contents of the memory address contains the low order bits of the effective address
// The next memory location contains the high order bits
func ind(cpu *CPU) uint16 {
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
func inx(cpu *CPU) uint16 {
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
func iny(cpu *CPU) uint16 {
    location := uint16(cpu.read(cpu.PC))
    cpu.PC++

    low := uint16(cpu.read(location))
    high := uint16(cpu.read(location+1)) & 0x00FF

    address := (high<<8) | low
    address += uint16(cpu.Y)

    return address
}

