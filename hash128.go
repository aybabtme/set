package set

import (
	"github.com/dgryski/go-farm"
	"github.com/dgryski/go-spooky"
)

var (
	hash128IsMutable MutableSet = NewHashFunc128(0, farm.Hash128, true)
)

// NewSpooky128 is a Hash128 with spooky hash for hasher.
func NewSpooky128(n int, collidePanics bool) *Hash128 {
	return NewHashFunc128(n, func(b []byte) (lo, hi uint64) { spooky.Hash128(b, &lo, &hi); return }, collidePanics)
}

// NewFarm128 is a Hash128 with farmhash for hasher.
func NewFarm128(n int, collidePanics bool) *Hash128 {
	return NewHashFunc128(n, farm.Hash128, collidePanics)
}

type uint128 struct{ lo, hi uint64 }

// Hash128 is a hash based set, using a 128 bits hasher.
type Hash128 struct {
	m             map[uint128]struct{}
	collidePanics bool
	fh128         func(s []byte) (uint64, uint64)
}

// NewHashFunc128 creates a hash set using a hash128 hasher func.
func NewHashFunc128(n int, fh128 func(s []byte) (uint64, uint64), collidePanics bool) *Hash128 {
	return &Hash128{
		m:             make(map[uint128]struct{}, n),
		collidePanics: collidePanics,
		fh128:         fh128,
	}
}

func (m *Hash128) get128Block(s string) uint128 {
	lo, hi := m.fh128([]byte(s))
	return uint128{lo: lo, hi: hi}
}

// Add the key to the set.
func (m *Hash128) Add(s string) {
	block := m.get128Block(s)
	if m.collidePanics {
		_, ok := m.m[block]
		if ok {
			panic("Collision with '" + s + "'")
		}
	}
	m.m[block] = q
}

// Contains tells if this key was in the set at least once.
func (m *Hash128) Contains(s string) bool { _, ok := m.m[m.get128Block(s)]; return ok }

// IsEmpty tells if this set is empty.
func (m *Hash128) IsEmpty() bool { return len(m.m) == 0 }

// Len is the length of this set.
func (m *Hash128) Len() int { return len(m.m) }

// Delete the element form this set.
func (m *Hash128) Delete(s string) { delete(m.m, m.get128Block(s)) }
