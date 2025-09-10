package cpu68000

import (
	"cpu-emulator/pkg/cpu"
	"cpu-emulator/pkg/memory"
	"fmt"
)

// CPU68000 implements the cpu.CPU interface for the Motorola 68000 processor.
type CPU68000 struct {
	D      [8]uint32 // Data registers D0-D7
	A      [8]uint32 // Address registers A0-A7 (A7 is SP)
	PC     uint32
	SR     uint16
	Memory *memory.Memory
}

func New(memorySize int) cpu.CPU {
	return &CPU68000{
		Memory: memory.New(memorySize),
	}
}

func (c *CPU68000) Reset() {
	// Typically, 68000 loads SP and PC from 0x0 and 0x4 at reset
	c.A[7] = uint32(c.Memory.Read(0x0))<<24 | uint32(c.Memory.Read(0x1))<<16 | uint32(c.Memory.Read(0x2))<<8 | uint32(c.Memory.Read(0x3))
	c.PC = uint32(c.Memory.Read(0x4))<<24 | uint32(c.Memory.Read(0x5))<<16 | uint32(c.Memory.Read(0x6))<<8 | uint32(c.Memory.Read(0x7))
	c.SR = 0x2700 // Supervisor mode, interrupts masked
}

func (c *CPU68000) Step() error {
	opcode := uint16(c.Memory.Read(uint16(c.PC))<<8 | c.Memory.Read(uint16(c.PC+1)))
	instr := InstructionSet[opcode]
	if instr.Execute == nil {
		return fmt.Errorf("unknown opcode: 0x%04X at PC: 0x%08X", opcode, c.PC)
	}
	instr.Execute(c)
	return nil
}

func (c *CPU68000) LoadProgram(program []uint8, startAddress uint16) {
	for i, value := range program {
		c.Memory.Write(uint16(startAddress)+uint16(i), value)
	}
}
