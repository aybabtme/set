package set

// Guarantees the implementation of those interfaces
var (
	MapSetIsMutable  MutableSet = make(GoMapSet)
	MapSetIsListable ListSet    = make(GoMapSet)
	q                           = struct{}{}
)

// GoMapSet is a set of string implemented using Go maps.
type GoMapSet map[string]struct{}

// NewGoMapSet creates a GoMapSet of capacity n.
func NewGoMapSet(n int) GoMapSet { return make(GoMapSet, n) }

// Add the key to the set.
func (m GoMapSet) Add(s string) { m[s] = q }

// Contains tells if this key was in the set at least once.
func (m GoMapSet) Contains(s string) bool { _, ok := m[s]; return ok }

// IsEmpty tells if this set is empty.
func (m GoMapSet) IsEmpty() bool { return len(m) == 0 }

// Len is the length of this set.
func (m GoMapSet) Len() int { return len(m) }

// Delete the element form this set.
func (m GoMapSet) Delete(s string) { delete(m, s) }

// Keys gives all the keys in this GoMapSet.
func (m GoMapSet) Keys() []string {
	keys := make([]string, m.Len())
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
