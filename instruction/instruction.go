package instruction

import (
	"errors"
	"fmt"
)

type Command byte

type Target byte

const (
	NOP Command = iota
	ADD
	INC
)

const (
	A Target = iota
	B
	C
	D
	E
	H
	L
)


type Instruction struct {
	Opcode byte
	Inst   Command
	Target1 Target
	Target2	Target
	Length int
	Cycles int
}

func From_opcode(opcode byte, prefixed bool) (Instruction, error) {
	if prefixed {
		return from_opcode_prefixed(opcode)
	} else {
		return from_opcode_not_prefixed(opcode)
	}
}

func from_opcode_not_prefixed(opcode byte) (Instruction, error) {
	res := Instruction{Opcode: opcode}
	switch opcode {
	// case 0x00:
	// 	// NOP
	// 	res.Inst = NOP
	case 0x3C:
		res.Inst = INC
		res.Target1 = A
		res.Length = 1
		res.Cycles = 4
	case 0x80:
		// ADD A, B
		res.Inst = ADD
		res.Target1 = A
		res.Target2 = B
		res.Length = 1
		res.Cycles = 4
	case 0x81:
		res.Inst = ADD
		res.Target1 = A
		res.Target2 = C
		res.Length = 1
		res.Cycles = 4
	case 0x82:
		res.Inst = ADD
		res.Target1 = A
		res.Target2 = D
		res.Length = 1
		res.Cycles = 4
	case 0x83:
		res.Inst = ADD
		res.Target1 = A
		res.Target2 = E
		res.Length = 1
		res.Cycles = 4
	case 0x84:
		res.Inst = ADD
		res.Target1 = A
		res.Target2 = H
		res.Length = 1
		res.Cycles = 4
	case 0x85:
		res.Inst = ADD
		res.Target1 = A
		res.Target2 = L
		res.Length = 1
		res.Cycles = 4
	case 0x87:
		// ADD A, A
		res.Inst = ADD
		res.Target1 = A
		res.Target2 = A
		res.Length = 1
		res.Cycles = 4
	default:

		return res, errors.New(fmt.Sprintf("Opcode Not implemented 0x%02X\n", opcode))
	}

	return res, nil
}

func from_opcode_prefixed(opcode byte) (Instruction, error) {
	res := Instruction{Opcode: opcode}
	switch opcode {
	default:
		return res, errors.New(fmt.Sprintf("Prefixed Opcode Not implemented 0x%02X\n", opcode))
	}
	//return res, nil
}

func (inst Instruction) Print() {
	fmt.Printf("Opcode: 0x%02X\n", inst.Opcode)
}