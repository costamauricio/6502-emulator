package bus

import (
	"encoding/hex"
	"strings"
)

type Bus struct {
	ram [64 * 1024]byte
}

func (bus *Bus) String() string {
    return hex.EncodeToString(bus.ram[:])
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

	return 0x00
}

func (bus *Bus) LoadRamFromString(memory string, offset uint16) error {
	encodedString := strings.ReplaceAll(memory, " ", "")

	decoded, err := hex.DecodeString(encodedString)
	if err != nil {
		return err
	}

	for index, content := range decoded {
		bus.ram[offset+uint16(index)] = content
	}

	return nil
}
