package cpu6502

import (
	"cpu-emulator/pkg/cpu"
	"cpu-emulator/pkg/memory"
	"fmt"
)

// CPU6502 implements the cpu.CPU interface for the 6502 processor.
type CPU6502 struct {
	A, X, Y, SP uint8
	PC          uint16
	Status      uint8
	Memory      *memory.Memory
}

// New returns a new 6502 CPU with the given memory size.
func New(memorySize int) cpu.CPU {
	return &CPU6502{
		SP:     0xFD,
		Memory: memory.New(memorySize),
	}
}

// Reset resets the CPU state.
func (c *CPU6502) Reset() {
	c.A = 0
	c.X = 0
	c.Y = 0
	c.SP = 0xFD
	lo := c.Memory.Read(0xFFFC)
	hi := c.Memory.Read(0xFFFD)
	c.PC = uint16(lo) | (uint16(hi) << 8)
	c.Status = 0x34
}

// Step fetches and executes the next instruction.
func (c *CPU6502) Step() error {
	opcode := c.Memory.Read(c.PC)
	instr := InstructionSet[opcode]
	if instr.Execute == nil {
		return fmt.Errorf("unknown opcode: 0x%02X at PC: 0x%04X", opcode, c.PC)
	}
	instr.Execute(c)
	return nil
}

// LoadProgram loads a program into memory at the given address.
func (c *CPU6502) LoadProgram(program []uint8, startAddress uint16) {
	for i, value := range program {
		c.Memory.Write(startAddress+uint16(i), value)
	}
}

// Flag helpers
func (c *CPU6502) setZeroFlag(v uint8) {
	if v == 0 {
		c.Status |= 0x02
	} else {
		c.Status &^= 0x02
	}
}
func (c *CPU6502) setNegativeFlag(v uint8) {
	if v&0x80 != 0 {
		c.Status |= 0x80
	} else {
		c.Status &^= 0x80
	}
}

// Addressing mode helpers
func (c *CPU6502) fetchImmediate() uint8 {
	return c.Memory.Read(c.PC + 1)
}
func (c *CPU6502) fetchZeroPage() uint8 {
	addr := c.Memory.Read(c.PC + 1)
	return c.Memory.Read(uint16(addr))
}
