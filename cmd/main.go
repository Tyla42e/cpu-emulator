package main

import (
	"cpu-emulator/pkg/cpu6502"
)

func main() {
	cpu := cpu6502.New(65536)
	// Example program: LDA #$42; TAX; BRK
	program := []uint8{0xA9, 0x42, 0xAA, 0x00}
	cpu.LoadProgram(program, 0x8000)
	// Set reset vector to 0x8000
	cpu.(*cpu6502.CPU6502).Memory.Write(0xFFFC, 0x00)
	cpu.(*cpu6502.CPU6502).Memory.Write(0xFFFD, 0x80)
	cpu.Reset()
	for i := 0; i < 3; i++ {
		cpu.Step()
	}
	// Inspect cpu.(*cpu6502.CPU6502).A, .X, .PC, etc.
}
