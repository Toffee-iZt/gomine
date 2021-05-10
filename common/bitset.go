package common

// Bits ...
type Bits uint

// Set ...
func (b *Bits) Set(bit Bits) {
	*b |= bit
}

// Clear ...
func (b *Bits) Clear(bit Bits) {
	*b &^= bit
}

// Toggle ...
func (b *Bits) Toggle(bit Bits) {
	*b ^= bit
}

// Has ...
func (b Bits) Has(bit Bits) bool {
	return b&bit != 0
}

// NewBitSet creates new bitset
func NewBitSet(bb ...Bits) Bits {
	var b Bits
	for i := range bb {
		b.Set(bb[i])
	}
	return b
}
