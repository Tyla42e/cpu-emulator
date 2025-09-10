package memory

// Memory is a simple byte-addressable memory model.
type Memory struct {
	Data []uint8
}

// New returns a new Memory instance of the given size.
func New(size int) *Memory {
	return &Memory{
		Data: make([]uint8, size),
	}
}

// Read returns the byte at the given address.
func (m *Memory) Read(addr uint16) uint8 {
	return m.Data[addr]
}

// Write sets the byte at the given address.
func (m *Memory) Write(addr uint16, value uint8) {
	m.Data[addr] = value
}
