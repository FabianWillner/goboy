package cpu

import (
	"errors"
	"fmt"

	"github.com/FabianWillner/goboy/instruction"
)
type CPU struct {
	Register registers
}

func New() CPU {
	c := CPU{}
	return c
}


func (c *CPU) Execute(opcode byte) error {
	// TODO: Build Instruction
	inst, err := instruction.New(opcode)
	if err != nil {
		return err
	}

	// TODO: execute instruction
	err = c.handleInstruction(inst)
	if err != nil {
		return err
	}

	return nil
}

func (c *CPU) GetReg(target instruction.Target) (byte, error) {
	switch target {
	case instruction.A:
		return c.Register.A, nil
	case instruction.B:
		return c.Register.B, nil
	case instruction.C:
		return c.Register.C, nil
	case instruction.D:
		return c.Register.D, nil
	case instruction.E:
		return c.Register.E, nil
	default:
		return 0, errors.New(fmt.Sprintf("Get Not Implemented CPU doesn't have register: %x\n", target))
	}
}

func (c *CPU) SetReg(target instruction.Target, val byte) (error) {
	switch target {
	case instruction.A:
		c.Register.A = val
		return nil
	case instruction.B:
		c.Register.B = val
		return nil
	case instruction.C:
		c.Register.C = val
		return  nil
	case instruction.D:
		c.Register.D = val
		return  nil
	case instruction.E:
		c.Register.E = val
		return  nil
	default:
		return errors.New(fmt.Sprintf("Set Not Implemented CPU doesn't have register: %x\n", target))
	}
}


func (c CPU) PrintRegisters(){
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