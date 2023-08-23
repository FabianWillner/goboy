package memory_bus

type MemoryBus struct {
	memory [0xFFFF + 1]byte
}

func (m *MemoryBus) Read_byte(address uint16) byte {
	return m.memory[address]
}

func (m *MemoryBus) Set_byte(address uint16, val byte) {
	m.memory[address] = val
}

func (m *MemoryBus) Mem_copy(offset int, data []byte) {
	copy(m.memory[offset:], data)
}