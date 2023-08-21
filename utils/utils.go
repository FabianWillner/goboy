package utils

import (
	"fmt"
)

func PrintByte(b byte) {
	fmt.Printf("Byte value: 0x%02X\n", b)
}

func PrintUint16(b uint16) {
	fmt.Printf("uint16 value: 0x%04X\n", b)
}