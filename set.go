package set

// Set answers question of the type: is this string a member?
type Set interface {
	Add(string)
	Contains(string) bool
	IsEmpty() bool
	Len() int
}

// MutableSet is a Set from which you can remove keys.
type MutableSet interface {
	Set
	Delete(string)
}

// ListSet is a Set that can return the its keys.
type ListSet interface {
	Set
	Keys() []string
}

// Union of the two list set, the result stored in the
// out set. Everything in A or (inclusive) B is the result.
func Union(a, b ListSet, out Set) {
	for _, k := range a.Keys() {
		out.Add(k)
	}
	for _, k := range b.Keys() {
		out.Add(k)
	}
}

// Intersect of the two list set, the result stored in
// the out set. Everything in both A and B is the result.
func Intersect(a ListSet, b, out Set) {
	for _, k := range a.Keys() {
		if b.Contains(k) {
			out.Add(k)
		}
	}
}

// Difference between a and b. The result is stored in
// the out set. Everything that is in A but not in B will
// be the result.
func Difference(a ListSet, b, out Set) {
	for _, k := range a.Keys() {
		if !b.Contains(k) {
			out.Add(k)
		}
	}
}

// XOR of the two list set, the result stored in the out
// set. Everything not in both A and B is the result.
func XOR(a, b ListSet, out Set) {
	for _, k := range a.Keys() {
		if !b.Contains(k) {
			out.Add(k)
		}
	}
	for _, k := range b.Keys() {
		if !a.Contains(k) {
			out.Add(k)
		}
	}
}
