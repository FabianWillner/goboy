package main

import (
	"fmt"

	"github.com/FabianWillner/goboy/cpu"
)



func main() {
	cpu := cpu.New()
	cpu.Register.A = 0xF0

	cpu.PrintRegisters()

	err := cpu.Execute(0x87)
	if err != nil {
		fmt.Println(err)
		return
	}

	cpu.PrintRegisters()

	err = cpu.Execute(0x88)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpu.Register.SetF(0b10100000)

	cpu.PrintRegisters()
	// fmt.Println("Hello, World!")
	// rom_filepath := "./testroms/cpu_instrs.gb"
	// file, err := os.Open(rom_filepath)
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// defer file.Close()

	// // Read the file content into a byte slice
	// fileBytes, err := io.ReadAll(file)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }

	// PrintByte(fileBytes[0])
}
