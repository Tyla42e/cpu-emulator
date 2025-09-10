package cpu6502

import "cpu-emulator/pkg/cpu"

// InstructionSet is the 6502 instruction set as an array for fast lookup.
var InstructionSet [256]cpu.Instruction

func init() {
	InstructionSet[0xA9] = cpu.Instruction{
		Name:           "LDA",
		Opcode:         0xA9,
		Cycles:         2,
		AddressingMode: cpu.Immediate,
		Execute: func(c cpu.CPU) {
			cpu := c.(*CPU6502)
			value := cpu.fetchImmediate()
			cpu.A = value
			cpu.setZeroFlag(value)
			cpu.setNegativeFlag(value)
			cpu.PC += 2
		},
	}
	InstructionSet[0xA5] = cpu.Instruction{
		Name:           "LDA",
		Opcode:         0xA5,
		Cycles:         3,
		AddressingMode: cpu.ZeroPage,
		Execute: func(c cpu.CPU) {
			cpu := c.(*CPU6502)
			value := cpu.fetchZeroPage()
			cpu.A = value
			cpu.setZeroFlag(value)
			cpu.setNegativeFlag(value)
			cpu.PC += 2
		},
	}
	InstructionSet[0xAA] = cpu.Instruction{
		Name:           "TAX",
		Opcode:         0xAA,
		Cycles:         2,
		AddressingMode: cpu.Implied,
		Execute: func(c cpu.CPU) {
			cpu := c.(*CPU6502)
			cpu.X = cpu.A
			cpu.setZeroFlag(cpu.X)
			cpu.setNegativeFlag(cpu.X)
			cpu.PC++
		},
	}
	// ...add more instructions as needed
}
