package set

import (
	"github.com/dgryski/dgohash"
	"github.com/dgryski/go-farm"
	"github.com/dgryski/go-spooky"
	"hash"
	"hash/adler32"
)

var (
	Hash32IsMutable MutableSet = NewHash32(0, dgohash.NewMurmur3_x86_32())
)

// NewSpooky32 is a Hash32 with spooky hash for hasher.
func NewSpooky32(n int) *Hash32 { return NewHashFunc32(n, spooky.Hash32) }

// NewFarm32 is a Hash32 with farmhash for hasher.
func NewFarm32(n int) *Hash32 { return NewHashFunc32(n, farm.Hash32) }

// NewAdler32 is a Hash32 with adler32 for hasher.
func NewAdler32(n int) *Hash32 { return NewHash32(n, adler32.New()) }

// NewMurmur32 is a Hash32 with murmur3 for hasher.
func NewMurmur32(n int) *Hash32 { return NewHash32(n, dgohash.NewMurmur3_x86_32()) }

// NewDjb32 is a Hash32 with djb32 for hasher.
func NewDjb32(n int) *Hash32 { return NewHash32(n, dgohash.NewDjb32()) }

// NewElf32 is a Hash32 with Elf32 for hasher.
func NewElf32(n int) *Hash32 { return NewHash32(n, dgohash.NewElf32()) }

// NewJava32 is a Hash32 with Java32 for hasher.
func NewJava32(n int) *Hash32 { return NewHash32(n, dgohash.NewJava32()) }

// NewJenkins32 is a Hash32 with Jenkins32 for hasher.
func NewJenkins32(n int) *Hash32 { return NewHash32(n, dgohash.NewJenkins32()) }

// NewSDBM32 is a Hash32 with SDBM32 for hasher.
func NewSDBM32(n int) *Hash32 { return NewHash32(n, dgohash.NewSDBM32()) }

// NewSQLite32 is a Hash32 with SQLite32 for hasher.
func NewSQLite32(n int) *Hash32 { return NewHash32(n, dgohash.NewSQLite32()) }

// NewSuperFastHash is a Hash32 with SuperFastHash for hasher.
func NewSuperFastHash(n int) *Hash32 { return NewHash32(n, dgohash.NewSuperFastHash()) }

func fromHash32(h32 hash.Hash32) func([]byte) uint32 {
	return func(b []byte) uint32 {
		h32.Reset()
		_, _ = h32.Write(b)
		return h32.Sum32()
	}
}

// Hash32 is a hash based set, using a 32 bits hasher.
type Hash32 struct {
	m    map[uint32]struct{}
	fh32 func(s []byte) uint32
}

// NewHash32 creates a hash set using a hash32 hasher.
func NewHash32(n int, h hash.Hash32) *Hash32 {
	return &Hash32{
		m:    make(map[uint32]struct{}, n),
		fh32: fromHash32(h),
	}
}

// NewHashFunc32 creates a hash set using a hash32 hasher.
func NewHashFunc32(n int, fh32 func(s []byte) uint32) *Hash32 {
	return &Hash32{
		m:    make(map[uint32]struct{}, n),
		fh32: fh32,
	}
}

func (m *Hash32) get32Block(s string) uint32 {
	return m.fh32([]byte(s))
}

// Add the key to the set.
func (m *Hash32) Add(s string) { m.m[m.get32Block(s)] = q }

// Contains tells if this key was in the set at least once.
func (m *Hash32) Contains(s string) bool { _, ok := m.m[m.get32Block(s)]; return ok }

// IsEmpty tells if this set is empty.
func (m *Hash32) IsEmpty() bool { return len(m.m) == 0 }

// Len is the length of this set.
func (m *Hash32) Len() int { return len(m.m) }

// Delete the element form this set.
func (m *Hash32) Delete(s string) { delete(m.m, m.get32Block(s)) }
