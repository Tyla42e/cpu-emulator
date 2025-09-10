package cpuz80

import "cpu-emulator/pkg/cpu"

var InstructionSet [256]cpu.Instruction

func init() {
	InstructionSet[0x3E] = cpu.Instruction{
		Name:           "LD A, n",
		Opcode:         0x3E,
		Cycles:         7,
		AddressingMode: cpu.Immediate,
		Execute: func(c cpu.CPU) {
			cpu := c.(*CPUZ80)
			value := cpu.Memory.Read(cpu.PC + 1)
			cpu.A = value
			cpu.PC += 2
		},
	}
	// ...add more instructions as needed
}
