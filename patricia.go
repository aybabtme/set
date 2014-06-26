package set

import (
	"github.com/tchap/go-patricia/patricia"
)

// Guarantees the implementation of those interfaces
var (
	tchapPatriciaIsMutable  MutableSet = NewTchapPatricia()
	tchapPatriciaIsListable ListSet    = NewTchapPatricia()
)

// TchapPatricia is a set of string implemented using a sorted slice of strings.
type TchapPatricia struct {
	p *patricia.Trie
}

// NewTchapPatricia creates a TchapPatricia of capacity n.
func NewTchapPatricia() *TchapPatricia {
	return &TchapPatricia{
		p: patricia.NewTrie(),
	}
}

// Add the key to the set.
func (pat *TchapPatricia) Add(s string) {
	pat.p.Set(patricia.Prefix(s), q)
}

// Contains tells if this key was in the set at least once.
func (pat TchapPatricia) Contains(s string) bool {
	return pat.p.Match(patricia.Prefix(s))
}

// Delete the element form this set.
func (pat *TchapPatricia) Delete(s string) {
	pat.p.Delete(patricia.Prefix(s))
}

// IsEmpty tells if this set is empty.
func (pat TchapPatricia) IsEmpty() bool {
	return pat.Len() == 0
}

// Len is the length of this set.
func (pat TchapPatricia) Len() (count int) {
	pat.p.Visit(func(_ patricia.Prefix, _ patricia.Item) error { count++; return nil })
	return count
}

// Keys gives all the keys in this TchapPatricia.
func (pat TchapPatricia) Keys() (keys []string) {
	pat.p.Visit(func(keyb patricia.Prefix, _ patricia.Item) error {
		keys = append(keys, string(keyb))
		return nil
	})
	return
}
