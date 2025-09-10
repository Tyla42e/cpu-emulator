package cpu6510

import "cpu-emulator/pkg/cpu"

var InstructionSet [256]cpu.Instruction

func init() {
	InstructionSet[0xA9] = cpu.Instruction{
		Name:           "LDA",
		Opcode:         0xA9,
		Cycles:         2,
		AddressingMode: cpu.Immediate,
		Execute: func(c cpu.CPU) {
			cpu := c.(*CPU6510)
			value := cpu.fetchImmediate()
			cpu.A = value
			cpu.PC += 2
		},
	}
	// ...add more instructions as needed
}
