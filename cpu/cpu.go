package cpu

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/FabianWillner/goboy/instruction"
	"github.com/FabianWillner/goboy/memory_bus"
)
type CPU struct {
	Register registers
	Bus memory_bus.MemoryBus
	PC uint16
	SP uint16
}

func New() CPU {
	c := CPU{Bus: memory_bus.MemoryBus{}}
	return c
}

func (c *CPU) PowerUpSequence() {
	c.Register.SetBC(0x0013)
	c.Register.SetDE(0x00D8)
	c.Register.SetHL(0x014D)
	c.SP = 0xFFFE
	c.Bus.Set_byte(0xFF05, 0x00)
	c.Bus.Set_byte(0xFF06, 0x00)
	c.Bus.Set_byte(0xFF07, 0x00)
	c.Bus.Set_byte(0xFF10, 0x80)
	c.Bus.Set_byte(0xFF11, 0xBF)
	c.Bus.Set_byte(0xFF12, 0xF3)
	c.Bus.Set_byte(0xFF14, 0xBF)
	c.Bus.Set_byte(0xFF16, 0x3F)
	c.Bus.Set_byte(0xFF17, 0x00)
	c.Bus.Set_byte(0xFF19, 0xBF)
	c.Bus.Set_byte(0xFF1A, 0x7F)
	c.Bus.Set_byte(0xFF1B, 0xFF)
	c.Bus.Set_byte(0xFF1C, 0x9F)
	c.Bus.Set_byte(0xFF1E, 0xBF)
	c.Bus.Set_byte(0xFF20, 0xFF)
	c.Bus.Set_byte(0xFF21, 0x00)
	c.Bus.Set_byte(0xFF22, 0x00)
	c.Bus.Set_byte(0xFF23, 0xBF)
	c.Bus.Set_byte(0xFF24, 0x77)
	c.Bus.Set_byte(0xFF25, 0xF3)
	c.Bus.Set_byte(0xFF26, 0xF1)
	c.Bus.Set_byte(0xFF40, 0x91)
	c.Bus.Set_byte(0xFF42, 0x00)
	c.Bus.Set_byte(0xFF43, 0x00)
	c.Bus.Set_byte(0xFF45, 0x00)
	c.Bus.Set_byte(0xFF47, 0xFC)
	c.Bus.Set_byte(0xFF48, 0xFF)
	c.Bus.Set_byte(0xFF49, 0xFF)
	c.Bus.Set_byte(0xFF4A, 0x00)
	c.Bus.Set_byte(0xFF4B, 0x00)
	c.Bus.Set_byte(0xFFFF, 0x00)
	c.PC = 0x0100
}

func (c *CPU) LoadRom(rom_filepath string) error {
	//rom_filepath := "./testroms/cpu_instrs.gb"
	file, err := os.Open(rom_filepath)
	if err != nil {
		return errors.New(fmt.Sprintf("Error opening file: %s", err))
	}
	defer file.Close()

	// Read the file content into a byte slice
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return errors.New(fmt.Sprintf("Error reading file: %s", err))
	}
	c.Bus.Mem_copy(0x0100, fileBytes)

	return nil
}

func (c *CPU) Step() error{
	instruction_byte := c.Bus.Read_byte(c.PC)
	prefixed := instruction_byte == 0xCB
	if prefixed {
		instruction_byte = c.Bus.Read_byte(c.PC + 1)
	}

	err := c.Execute(instruction_byte, prefixed)
	if err != nil {
		return err
	}

	return nil

}


func (c *CPU) Execute(opcode byte, prefixed bool)  error {
	// TODO: Build Instruction
	inst, err := instruction.From_opcode(opcode, prefixed)
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