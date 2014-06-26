package set

import (
	"sort"
)

// Guarantees the implementation of those interfaces
var (
	SortedSetIsMutable  MutableSet = NewSortedSet(0)
	SortedSetIsListable ListSet    = NewSortedSet(0)
)

// SortedSet is a set of string implemented using a sorted slice of strings.
type SortedSet struct {
	s []string
}

// NewSortedSet creates a SortedSet of capacity n.
func NewSortedSet(n int) *SortedSet {
	return &SortedSet{
		s: make([]string, 0, n),
	}
}

// Add the key to the set.
func (sset *SortedSet) Add(s string) {
	sset.s = append(sset.s, s)
	sort.Strings(sset.s)
}

// Contains tells if this key was in the set at least once.
func (sset SortedSet) Contains(s string) bool {
	i := sort.SearchStrings(sset.s, s)
	return i < len(sset.s) && sset.s[i] == s
}

// Delete the element form this set.
func (sset *SortedSet) Delete(s string) {
	i := sort.SearchStrings(sset.s, s)
	if i >= len(sset.s) {
		// not in this set
		return
	}
	sset.s[i] = sset.s[len(sset.s)-1]
	sset.s = sset.s[:len(sset.s)-1]
	if i != len(sset.s)-1 {
		sort.Strings(sset.s)
	}
}

// IsEmpty tells if this set is empty.
func (sset SortedSet) IsEmpty() bool { return len(sset.s) == 0 }

// Len is the length of this set.
func (sset SortedSet) Len() int { return len(sset.s) }

// Keys gives all the keys in this SortedSet.
func (sset SortedSet) Keys() []string { return sset.s }
