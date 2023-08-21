package utils

import (
	"fmt"

	"github.com/FabianWillner/goboy/cpu"
)

func PrintByte(b byte) {
	fmt.Printf("Byte value: 0x%02X\n", b)
}

func PrintUint16(b uint16) {
	fmt.Printf("uint16 value: 0x%04X\n", b)
}

func PrintRegisters(c cpu.CPU){
	reg := c.Register
	fmt.Println("-------------------")
	fmt.Printf("A: 0x%02X\tF: 0x%02X\n", reg.A, reg.F())
	fmt.Printf("B: 0x%02X\tC: 0x%02X\n", reg.B, reg.C)
	fmt.Printf("D: 0x%02X\tE: 0x%02X\n", reg.D, reg.E)
	fmt.Printf("H: 0x%02X\tL: 0x%02X\n", reg.H, reg.L)
	fmt.Printf("AF: 0x%04X BC: 0x%04X\n", reg.GetAF(), reg.GetBC())
	fmt.Printf("DE: 0x%04X HL: 0x%04X\n", reg.GetDE(), reg.GetHL())
	fmt.Printf("Zero: %d\t\tSubtract: %d\n", boolToInt(reg.Flag.Zero), boolToInt(reg.Flag.Subtract))
	fmt.Printf("Half Carry: %d\tCarry: %d\n", boolToInt(reg.Flag.Half_carry), boolToInt(reg.Flag.Carry))
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}