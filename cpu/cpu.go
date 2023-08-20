package cpu

type Flags struct {
	Zero       bool
	Subtract   bool
	Half_carry bool
	Carry      bool
}

func (f Flags) GetF() byte {
	res := byte(0)

	if f.Zero {
		res = (1 << 7) | res
	}

	if f.Subtract {
		res = (1 << 6) | res
	}

	if f.Half_carry {
		res = (1 << 5) | res
	}

	if f.Carry {
		res = (1 << 4) | res
	}

	return res
}

func (f *Flags) SetF(val byte) {
	f.Zero = (val>>7)&0b1 != 0
	f.Subtract = (val>>6)&0b1 != 0
	f.Half_carry = (val>>5)&0b1 != 0
	f.Carry = (val>>4)&0b1 != 0
}

type CPU struct {
	A    byte
	B    byte
	C    byte
	D    byte
	E    byte
	H    byte
	L    byte
	Flag Flags
}

func New() CPU {
	c := CPU{}
	return c
}

func (c CPU) F() byte {
	return c.Flag.GetF()
}

func (c *CPU) SetF(val byte) {
	c.Flag.SetF(val)
}

func (c CPU) GetAF() uint16 {
	AF := (uint16(c.A) << 8) | uint16(c.F())
	return AF
}

func (c *CPU) SetAF(val uint16) {
	c.A = byte((val & 0xFF00) >> 8)
	c.SetF(byte((val & 0xFF)))
}

func (c CPU) GetBC() uint16 {
	BC := (uint16(c.B) << 8) | uint16(c.C)
	return BC
}

func (c *CPU) SetBC(val uint16) {
	c.B = byte((val & 0xFF00) >> 8)
	c.C = byte((val & 0xFF))
}

func (c CPU) GetDE() uint16 {
	DE := (uint16(c.D) << 8) | uint16(c.E)
	return DE
}

func (c *CPU) SetDE(val uint16) {
	c.D = byte((val & 0xFF00) >> 8)
	c.E = byte((val & 0xFF))
}

func (c CPU) GetHL() uint16 {
	HL := (uint16(c.H) << 8) | uint16(c.L)
	return HL
}

func (c *CPU) SetHL(val uint16) {
	c.H = byte((val & 0xFF00) >> 8)
	c.L = byte((val & 0xFF))
}