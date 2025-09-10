package cpu65c02

import (
	"cpu-emulator/pkg/cpu"
	"cpu-emulator/pkg/memory"
	"fmt"
)

// CPU65C02 implements the cpu.CPU interface for the WDC 65C02 processor.
type CPU65C02 struct {
	A, X, Y, SP uint8
	PC          uint16
	Status      uint8
	Memory      *memory.Memory
}

func New(memorySize int) cpu.CPU {
	return &CPU65C02{
		SP:     0xFF,
		Memory: memory.New(memorySize),
	}
}

func (c *CPU65C02) Reset() {
	c.A = 0
	c.X = 0
	c.Y = 0
	c.SP = 0xFF
	lo := c.Memory.Read(0xFFFC)
	hi := c.Memory.Read(0xFFFD)
	c.PC = uint16(lo) | (uint16(hi) << 8)
	c.Status = 0x34
}

func (c *CPU65C02) Step() error {
	opcode := c.Memory.Read(c.PC)
	instr := InstructionSet[opcode]
	if instr.Execute == nil {
		return fmt.Errorf("unknown opcode: 0x%02X at PC: 0x%04X", opcode, c.PC)
	}
	instr.Execute(c)
	return nil
}

func (c *CPU65C02) LoadProgram(program []uint8, startAddress uint16) {
	for i, value := range program {
		c.Memory.Write(startAddress+uint16(i), value)
	}
}
