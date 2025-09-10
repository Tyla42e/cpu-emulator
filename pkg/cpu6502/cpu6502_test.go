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
