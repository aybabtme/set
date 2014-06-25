package set

// Guarantees the implementation of those interfaces
var (
	GoMapIsMutable  MutableSet = NewGoMap(0)
	GoMapIsListable ListSet    = NewGoMap(0)
	q                          = struct{}{}
)

// GoMap is a set of string implemented using Go maps.
type GoMap map[string]struct{}

// NewGoMap creates a GoMap of capacity n.
func NewGoMap(n int) GoMap { return make(GoMap, n) }

// Add the key to the set.
func (m GoMap) Add(s string) { m[s] = q }

// Contains tells if this key was in the set at least once.
func (m GoMap) Contains(s string) bool { _, ok := m[s]; return ok }

// IsEmpty tells if this set is empty.
func (m GoMap) IsEmpty() bool { return len(m) == 0 }

// Len is the length of this set.
func (m GoMap) Len() int { return len(m) }

// Delete the element form this set.
func (m GoMap) Delete(s string) { delete(m, s) }

// Keys gives all the keys in this GoMap.
func (m GoMap) Keys() []string {
	keys := make([]string, m.Len())
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
