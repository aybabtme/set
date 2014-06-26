package set

import (
	"github.com/apokalyptik/quicktrie"
)

// Guarantees the implementation of those interfaces
var (
	QuicktrieIsMutable  MutableSet = NewQuicktrie()
	QuicktrieIsListable ListSet    = NewQuicktrie()
)

// Quicktrie is a set of string implemented using a sorted slice of strings.
type Quicktrie struct {
	r *trie.Trie
}

// NewQuicktrie creates a Quicktrie of capacity n.
func NewQuicktrie() *Quicktrie {
	return &Quicktrie{
		r: trie.NewBWTrie(),
	}
}

// Add the key to the set.
func (quick *Quicktrie) Add(s string) {
	quick.r.Add(s)
}

// Contains tells if this key was in the set at least once.
func (quick Quicktrie) Contains(s string) bool {
	return quick.r.Exists(s)
}

// IsEmpty tells if this set is empty.
func (quick Quicktrie) IsEmpty() bool {
	return quick.Len() == 0
}

// Delete the element form this set.
func (quick *Quicktrie) Delete(s string) {
	quick.r.Del(s)
}

// Len is the length of this set.
func (quick Quicktrie) Len() int {
	return quick.r.Count()
}

// Keys gives all the keys in this Quicktrie.
func (quick Quicktrie) Keys() (keys []string) {
	quick.r.Iterate(func(keyb []byte, _ interface{}) {
		keys = append(keys, string(keyb))
	})
	return
}
