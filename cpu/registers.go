package cpu

type registers struct {
	A    byte
	B    byte
	C    byte
	D    byte
	E    byte
	H    byte
	L    byte
	Flag Flags
}

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

func (reg registers) F() byte {
	return reg.Flag.GetF()
}

func (reg *registers) SetF(val byte) {
	reg.Flag.SetF(val)
}

func (reg registers) GetAF() uint16 {
	return (uint16(reg.A) << 8) | uint16(reg.F())
}

func (reg *registers) SetAF(val uint16) {
	reg.A = byte((val & 0xFF00) >> 8)
	reg.SetF(byte((val & 0xFF)))
}

func (reg registers) GetBC() uint16 {
	return (uint16(reg.B) << 8) | uint16(reg.C)
}

func (reg *registers) SetBC(val uint16) {
	reg.B = byte((val & 0xFF00) >> 8)
	reg.C = byte((val & 0xFF))
}

func (reg registers) GetDE() uint16 {
	return (uint16(reg.D) << 8) | uint16(reg.E)
}

func (reg *registers) SetDE(val uint16) {
	reg.D = byte((val & 0xFF00) >> 8)
	reg.E = byte((val & 0xFF))
}

func (reg registers) GetHL() uint16 {
	return (uint16(reg.H) << 8) | uint16(reg.L)
}

func (reg *registers) SetHL(val uint16) {
	reg.H = byte((val & 0xFF00) >> 8)
	reg.L = byte((val & 0xFF))
}