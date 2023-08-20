package main

import (
	"github.com/FabianWillner/goboy/cpu"
	"github.com/FabianWillner/goboy/utils"
)



func main() {
	val := byte(0xFF)
	gba_cpu := cpu.New()
	utils.PrintRegisters(gba_cpu)

	gba_cpu.A = val
	gba_cpu.E = val

	utils.PrintRegisters(gba_cpu)

	gba_cpu.SetF(0b10100000)

	utils.PrintRegisters(gba_cpu)
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
