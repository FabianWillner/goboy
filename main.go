package main

import (
	"fmt"

	"github.com/FabianWillner/goboy/cpu"
)



func main() {
	cpu := cpu.New()
	cpu.PowerUpSequence()
	//pu.PrintRegisters()
	err := cpu.LoadRom("./testroms/cpu_instrs.gb")
	if err != nil {
		fmt.Println(err)
		return
	}

	for true {
		err = cpu.Step()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	// for i := 0; i < 10; i++ {
	// 	utils.PrintByte(cpu.Bus.Read_byte(uint16(0x00FF + i)))
	// }
	
	fmt.Println("Run Complete!")
	
}
