package main

import "log"

func main() {
	nes := NewBus()
	nes.cpu.write(0, 0xA9)
	nes.cpu.write(1, 0x0C)
	nes.cpu.write(2, 0x0D)
	nes.cpu.write(0x0D, 0x0F)

	log.Println(nes.cpu.read(0))
	log.Println(nes.cpu.read(1))

	for {
		nes.cpu.Clock()

		if nes.cpu.cycle == 0 {
			break
		}
	}

	log.Printf("%v", nes.cpu.ac)
}
