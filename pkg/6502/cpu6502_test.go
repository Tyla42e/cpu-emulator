package cpu6502

import "testing"

func TestLDAImmediate(t *testing.T) {
	cpu := New(65536).(*CPU6502)
	cpu.Memory.Write(0xFFFC, 0x00)
	cpu.Memory.Write(0xFFFD, 0x80)
	cpu.Memory.Write(0x8000, 0xA9) // LDA #$42
	cpu.Memory.Write(0x8001, 0x42)
	cpu.Reset()
	if err := cpu.Step(); err != nil {
		t.Fatal(err)
	}
	if cpu.A != 0x42 {
		t.Errorf("Expected A=0x42, got 0x%02X", cpu.A)
	}
}

func TestADCImmediate_NoCarry(t *testing.T) {
	cpu := New(65536).(*CPU6502)
	// Set reset vector to 0x8000
	cpu.Memory.Write(0xFFFC, 0x00)
	cpu.Memory.Write(0xFFFD, 0x80)
	// Program: ADC #$10
	cpu.Memory.Write(0x8000, 0x69) // ADC Immediate opcode
	cpu.Memory.Write(0x8001, 0x10) // Operand
	cpu.A = 0x20
	cpu.Status &^= 0x01 // Clear carry flag
	cpu.Reset()
	cpu.PC = 0x8000 // Ensure PC is at start of program

	if err := cpu.Step(); err != nil {
		t.Fatal(err)
	}
	if cpu.A != 0x30 {
		t.Errorf("Expected A=0x30, got 0x%02X", cpu.A)
	}
	if cpu.Status&0x01 != 0 {
		t.Errorf("Carry flag should not be set")
	}
}

func TestADCImmediate_WithCarry(t *testing.T) {
	cpu := New(65536).(*CPU6502)
	cpu.Memory.Write(0xFFFC, 0x00)
	cpu.Memory.Write(0xFFFD, 0x80)
	cpu.Memory.Write(0x8000, 0x69)
	cpu.Memory.Write(0x8001, 0xFF)
	cpu.A = 0x02
	cpu.Status |= 0x01 // Set carry flag
	cpu.Reset()
	cpu.PC = 0x8000

	if err := cpu.Step(); err != nil {
		t.Fatal(err)
	}
	if cpu.A != 0x02 {
		t.Errorf("Expected A=0x02, got 0x%02X", cpu.A)
	}
	if cpu.Status&0x01 == 0 {
		t.Errorf("Carry flag should be set")
	}
}

func TestADCImmediate_Overflow(t *testing.T) {
	cpu := New(65536).(*CPU6502)
	cpu.Memory.Write(0xFFFC, 0x00)
	cpu.Memory.Write(0xFFFD, 0x80)
	cpu.Memory.Write(0x8000, 0x69)
	cpu.Memory.Write(0x8001, 0x50)
	cpu.A = 0x50
	cpu.Status &^= 0x01 // Clear carry flag
	cpu.Reset()
	cpu.PC = 0x8000

	if err := cpu.Step(); err != nil {
		t.Fatal(err)
	}
	if cpu.A != 0xA0 {
		t.Errorf("Expected A=0xA0, got 0x%02X", cpu.A)
	}
	if cpu.Status&0x40 == 0 {
		t.Errorf("Overflow flag should be set")
	}
	if cpu.Status&0x80 == 0 {
		t.Errorf("Negative flag should be set")
	}
}
