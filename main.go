package main

func main() {
	nes := NewBus()
	nes.Write(0, 0xA9)
	nes.Write(1, 0x0C)
	nes.Write(2, 0x0D)

	for {
		nes.cpu.Clock()

		if nes.cpu.cycle == 0 {
			break
		}
	}
}
