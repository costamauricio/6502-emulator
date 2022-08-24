package bus

type Bus struct {
    ram [64 *1024]byte
}

func (bus *Bus) Write(address uint16, data byte) {
    if address <= 0xFFFF {
        bus.ram[address] = data
    }
}

func (bus *Bus) Read(address uint16) byte {
    if address <= 0xFFFF {
        return bus.ram[address]
    }

    return 0x00;
}
