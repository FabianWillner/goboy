package cpu

import (
	"errors"
	"fmt"
	"math"

	"github.com/FabianWillner/goboy/instruction"
)

func (c *CPU) handleInstruction(inst instruction.Instruction) error {
	switch inst.Inst {
	case instruction.ADD:
		target1, err := c.GetReg(inst.Target1)
		if err != nil{
			return err
		}
		target2, err := c.GetReg(inst.Target2)
		if err != nil{
			return err
		}

		val, f := add(target1, target2)
		c.SetReg(inst.Target1, val)
		c.Register.SetF(f)
		c.PC += 1

	default:
		return errors.New(fmt.Sprintf("Not implemented 0x%02X, Instruction: %x\n", inst.Opcode, inst.Inst))
	}
	return nil
}

// returns the value and the F register
func add(a byte, b byte) (byte, byte){
	flag := Flags{}
	tmp := uint16(a) + uint16(b)
	result := byte(tmp)
	flag.Carry = tmp > math.MaxUint8
	flag.Subtract = false
	flag.Zero = result == 0
	flag.Half_carry = (a & 0xF) + (b & 0xF) > 0xF;
	return result, flag.GetF()
}

func inc(a byte) (byte, byte) {
	// flag := Flags{}
	// result := byte(a + 1)

}