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
}

func New(opcode byte) (Instruction, error) {
	res := Instruction{Opcode: opcode}
	switch opcode {
	// case 0x00:
	// 	// NOP
	// 	res.Inst = NOP
	case 0x80:
		// ADD A, B
		res.Inst = ADD
		res.Target1 = A
		res.Target2 = B
	case 0x81:
		res.Inst = ADD
		res.Target1 = A
		res.Target2 = C
	case 0x82:
		res.Inst = ADD
		res.Target1 = A
		res.Target2 = D
	case 0x83:
		res.Inst = ADD
		res.Target1 = A
		res.Target2 = E
	case 0x84:
		res.Inst = ADD
		res.Target1 = A
		res.Target2 = H
	case 0x85:
		res.Inst = ADD
		res.Target1 = A
		res.Target2 = L
	case 0x87:
		// ADD A, A
		res.Inst = ADD
		res.Target1 = A
		res.Target2 = A
	default:

		return res, errors.New(fmt.Sprintf("Opcode Not implemented 0x%02X\n", opcode))
	}

	return res, nil
}

func (inst Instruction) Print() {
	fmt.Printf("Opcode: 0x%02X\n", inst.Opcode)
}