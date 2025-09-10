package cpu6502

import "cpu-emulator/pkg/cpu"

// InstructionSet is the 6502 instruction set as an array for fast lookup.
var InstructionSet [256]cpu.Instruction

func (c *CPU6502) setCarryFlag(v bool) {
	if v {
		c.Status |= 0x01
	} else {
		c.Status &^= 0x01
	}
}
func (c *CPU6502) setOverflowFlag(v bool) {
	if v {
		c.Status |= 0x40
	} else {
		c.Status &^= 0x40
	}
}

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
	InstructionSet[0x69] = cpu.Instruction{
		Name:           "ADC",
		Opcode:         0x69,
		Cycles:         2,
		AddressingMode: cpu.Immediate,
		Execute: func(c cpu.CPU) {
			cpu := c.(*CPU6502)
			value := cpu.fetchImmediate()
			carry := uint16(0)
			if cpu.Status&0x01 != 0 {
				carry = 1
			}
			sum := uint16(cpu.A) + uint16(value) + carry

			// Set flags
			cpu.setCarryFlag(sum > 0xFF)
			result := uint8(sum)
			cpu.setZeroFlag(result)
			cpu.setNegativeFlag(result)
			// Overflow: if sign of A == sign of value, but sign of result != sign of A
			cpu.setOverflowFlag(((cpu.A ^ result) & (value ^ result) & 0x80) != 0)

			cpu.A = result
			cpu.PC += 2
		},
	}
	// ...add more instructions as needed
}
