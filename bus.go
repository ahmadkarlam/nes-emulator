package main

type Bus struct {
	cpu CPU
	ram [2048]uint8
}

func (b *Bus) Read(address uint16) uint8 {
	return b.ram[address]
}

func (b *Bus) Write(address uint16, data uint8) {
	if address >= 0x0000 && address <= 0x1FFF {
		b.ram[address] = data
	}
}

func NewBus() Bus {
	b := Bus{}
	b.cpu = NewCPU(&b)
	return b
}
