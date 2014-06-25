package set

import (
	"crypto/sha1"
)

type sha1block [sha1.Size]byte

var (
	HashSHA1IsMutable MutableSet = NewHashSHA1(0)
)

// HashSHA1 is a hash based set, using SHA1 for hashing.
type HashSHA1 struct {
	m map[sha1block]struct{}
}

// NewHashSHA1 creates a hash set using SHA1.
func NewHashSHA1(n int) *HashSHA1 {
	return &HashSHA1{m: make(map[sha1block]struct{}, n)}
}

func getSHA1Block(s string) sha1block { return sha1.Sum([]byte(s)) }

// Add the key to the set.
func (m *HashSHA1) Add(s string) { m.m[getSHA1Block(s)] = q }

// Contains tells if this key was in the set at least once.
func (m *HashSHA1) Contains(s string) bool { _, ok := m.m[getSHA1Block(s)]; return ok }

// IsEmpty tells if this set is empty.
func (m *HashSHA1) IsEmpty() bool { return len(m.m) == 0 }

// Len is the length of this set.
func (m *HashSHA1) Len() int { return len(m.m) }

// Delete the element form this set.
func (m *HashSHA1) Delete(s string) { delete(m.m, getSHA1Block(s)) }
