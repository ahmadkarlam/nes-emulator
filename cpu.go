package main

const (
	C uint8 = (1 << 0) // Carry Bit
	Z uint8 = (1 << 1) // Zero
	I uint8 = (1 << 2) // Disable Interrupts
	D uint8 = (1 << 3) // Decimal Mode (unused in this implementation)
	B uint8 = (1 << 4) // Break
	U uint8 = (1 << 5) // Unused
	V uint8 = (1 << 6) // Overflow
	N uint8 = (1 << 7) // Negative
)

type instructionSet struct {
	Name        string
	AddressMode func() uint8
	Instruction func() uint8
	Cycles      uint8
}

var lookup = [0xFF]instructionSet{}

type CPU struct {
	// program counter
	pc uint16
	// accumulator
	ac uint8
	// X Register
	x uint8
	// Y Register
	y uint8
	// status register
	// N	Negative
	// V	Overflow
	// -	ignored
	// B	Break
	// D	Decimal (use BCD for arithmetics)
	// I	Interrupt (IRQ disable)
	// Z	Zero
	// C	Carry
	sr uint8
	// stack pointer
	sp uint8
	// bus
	bus         *Bus
	cycle       uint8
	addressTemp uint16
}

func NewCPU(bus *Bus) *CPU {
	c := CPU{
		bus: bus,
	}
	mappingLookupInstructionSet(&c)
	return &c
}

func (c *CPU) setFlag(flag uint8, value bool) {
	if value {
		c.sr |= flag
	} else {
		c.sr &= ^flag
	}
}

func (c *CPU) Clock() {
	if c.cycle == 0 {
		c.doCycle()
	}

	c.cycle--
}

func (c *CPU) doCycle() {
	operationCode := c.bus.Read(c.pc)
	c.setFlag(U, true)

	c.pc++

	cycles := lookup[operationCode].Cycles

	cycles += lookup[operationCode].AddressMode()
	cycles += lookup[operationCode].Instruction()

	c.setFlag(U, true)
}

// Instruction Set
func (c *CPU) lda() uint8 {
	c.ac = c.bus.Read(c.addressTemp)
	c.setFlag(Z, c.ac == 0x00)
	c.setFlag(N, (c.ac&0x80) > 0)
	return 1
}

// Address mode
func (c *CPU) imp() uint8 {
	return 0
}

func (c *CPU) imm() uint8 {
	c.pc++
	c.addressTemp = c.pc
	return 0
}

func mappingLookupInstructionSet(c *CPU) {
	lookup[0xA9] = instructionSet{
		"LDA", c.imm, c.lda, 2,
	}
}
