package cpu68000

import "cpu-emulator/pkg/cpu"

type Instruction68000 struct {
	Name           string
	Opcode         uint16
	Cycles         int
	AddressingMode cpu.AddressingMode
	Execute        func(cpu *CPU68000)
}

var InstructionSet [65536]Instruction68000

func init() {
	InstructionSet[0x4E71] = Instruction68000{
		Name:           "NOP",
		Opcode:         0x4E71,
		Cycles:         4,
		AddressingMode: cpu.Implied,
		Execute: func(c *CPU68000) {
			c.PC += 2
		},
	}
	// Add more instructions as needed
}
