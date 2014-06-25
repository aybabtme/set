package set

import (
	"github.com/dgryski/go-farm"
	"github.com/dgryski/go-spooky"
	"hash"
)

var (
	hash64IsMutable MutableSet = NewHash64(0, nil)
)

// NewSpooky64 is a Hash64 with spooky hash for hasher.
func NewSpooky64(n int) *Hash64 { return NewHashFunc64(n, spooky.Hash64) }

// NewFarm64 is a Hash64 with farmhash for hasher.
func NewFarm64(n int) *Hash64 { return NewHashFunc64(n, farm.Hash64) }

func fromHash64(h64 hash.Hash64) func([]byte) uint64 {
	return func(b []byte) uint64 {
		h64.Reset()
		_, _ = h64.Write(b)
		return h64.Sum64()
	}
}

// Hash64 is a hash based set, using a 64 bits hasher.
type Hash64 struct {
	m    map[uint64]struct{}
	fh64 func(s []byte) uint64
}

// NewHash64 creates a hash set using a hash64 hasher.
func NewHash64(n int, h hash.Hash64) *Hash64 {
	return NewHashFunc64(n, fromHash64(h))
}

// NewHashFunc64 creates a hash set using a hash64 hasher func.
func NewHashFunc64(n int, fh64 func(s []byte) uint64) *Hash64 {
	return &Hash64{
		m:    make(map[uint64]struct{}, n),
		fh64: fh64,
	}
}

func (m *Hash64) get64Block(s string) uint64 {
	return m.fh64([]byte(s))
}

// Add the key to the set.
func (m *Hash64) Add(s string) { m.m[m.get64Block(s)] = q }

// Contains tells if this key was in the set at least once.
func (m *Hash64) Contains(s string) bool { _, ok := m.m[m.get64Block(s)]; return ok }

// IsEmpty tells if this set is empty.
func (m *Hash64) IsEmpty() bool { return len(m.m) == 0 }

// Len is the length of this set.
func (m *Hash64) Len() int { return len(m.m) }

// Delete the element form this set.
func (m *Hash64) Delete(s string) { delete(m.m, m.get64Block(s)) }
