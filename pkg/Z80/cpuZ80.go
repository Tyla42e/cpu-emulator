package cpuz80

import (
	"cpu-emulator/pkg/cpu"
	"cpu-emulator/pkg/memory"
	"fmt"
)

// CPUZ80 implements the cpu.CPU interface for the Z80 processor.
type CPUZ80 struct {
	A, F, B, C, D, E, H, L uint8
	SP, PC                 uint16
	Memory                 *memory.Memory
}

func New(memorySize int) cpu.CPU {
	return &CPUZ80{
		SP:     0xFFFF,
		PC:     0x0000,
		Memory: memory.New(memorySize),
	}
}

func (c *CPUZ80) Reset() {
	c.A, c.F, c.B, c.C, c.D, c.E, c.H, c.L = 0, 0, 0, 0, 0, 0, 0, 0
	c.SP = 0xFFFF
	c.PC = 0x0000
}

func (c *CPUZ80) Step() error {
	opcode := c.Memory.Read(c.PC)
	instr := InstructionSet[opcode]
	if instr.Execute == nil {
		return fmt.Errorf("unknown opcode: 0x%02X at PC: 0x%04X", opcode, c.PC)
	}
	instr.Execute(c)
	return nil
}

func (c *CPUZ80) LoadProgram(program []uint8, startAddress uint16) {
	for i, value := range program {
		c.Memory.Write(startAddress+uint16(i), value)
	}
}

// Flag helpers, addressing mode helpers as needed...
