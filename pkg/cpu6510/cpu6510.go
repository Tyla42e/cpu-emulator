package cpu6510

import (
	"cpu-emulator/pkg/cpu"
	"cpu-emulator/pkg/memory"
	"fmt"
)

// CPU6510 implements the cpu.CPU interface for the 6510 processor.
type CPU6510 struct {
	A, X, Y, SP uint8
	PC          uint16
	Memory      *memory.Memory
}

func New(memorySize int) cpu.CPU {
	return &CPU6510{
		SP:     0xFF,
		Memory: memory.New(memorySize),
	}
}

func (c *CPU6510) Reset() {
	c.A = 0
	c.X = 0
	c.Y = 0
	c.SP = 0xFF
	c.PC = 0x0000
}

func (c *CPU6510) Step() error {
	opcode := c.Memory.Read(c.PC)
	instr := InstructionSet[opcode]
	if instr.Execute == nil {
		return fmt.Errorf("unknown opcode: 0x%02X at PC: 0x%04X", opcode, c.PC)
	}
	instr.Execute(c)
	return nil
}

func (c *CPU6510) LoadProgram(program []uint8, startAddress uint16) {
	for i, value := range program {
		c.Memory.Write(startAddress+uint16(i), value)
	}
}

// Flag helpers
func (c *CPU6510) setZeroFlag(v uint8) {
	// ...implement as needed...
}
func (c *CPU6510) setNegativeFlag(v uint8) {
	// ...implement as needed...
}

// Addressing mode helpers
func (c *CPU6510) fetchImmediate() uint8 {
	return c.Memory.Read(c.PC + 1)
}
