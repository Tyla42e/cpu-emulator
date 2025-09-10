package cpu65c02

import (
	cpu6502 "cpu-emulator/pkg/6502"
	"cpu-emulator/pkg/cpu"
)

var InstructionSet [256]cpu.Instruction

func init() {
	// Step 1: Copy 6502 instructions
	InstructionSet = cpu6502.InstructionSet

	// Step 2: Override/add 65C02-specific instructions
	InstructionSet[0x12] = cpu.Instruction{
		Name:           "ORA (zp)",
		Opcode:         0x12,
		Cycles:         5,
		AddressingMode: cpu.ZeroPage, // Example, use correct mode
		Execute: func(c cpu.CPU) {
			// Implement 65C02-specific ORA (zp)
		},
	}
	InstructionSet[0xEA] = cpu.Instruction{
		Name:           "NOP",
		Opcode:         0xEA,
		Cycles:         2,
		AddressingMode: cpu.Implied,
		Execute: func(c cpu.CPU) {
			cpu := c.(*CPU65C02)
			cpu.PC++
		},
	}
	// ...add/override more as needed
}
