package cpu

// AddressingMode represents the CPU addressing mode.
type AddressingMode int

const (
	Implied AddressingMode = iota
	Immediate
	ZeroPage
	ZeroPageX
	ZeroPageY
	Absolute
	AbsoluteX
	AbsoluteY
	Indirect
	IndexedIndirect
	IndirectIndexed
	Relative
	Accumulator
)

// CPU is the interface all CPU implementations must satisfy.
type CPU interface {
	Reset()
	Step() error
	LoadProgram(program []uint8, startAddress uint16)
}

// Instruction represents a single CPU instruction.
type Instruction struct {
	Name           string
	Opcode         uint8
	Cycles         int
	AddressingMode AddressingMode
	Execute        func(cpu CPU)
}
