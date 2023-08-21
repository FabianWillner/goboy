package main

import (
	"github.com/FabianWillner/goboy/cpu"
	"github.com/FabianWillner/goboy/utils"
)



func main() {
	val := byte(0xFF)
	cpu := cpu.New()
	utils.PrintRegisters(cpu)

	cpu.Register.A = val
	cpu.Register.E = val

	utils.PrintRegisters(cpu)

	cpu.Register.SetF(0b10100000)

	utils.PrintRegisters(cpu)
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
