package set_test

import (
	"github.com/aybabtme/set"
	"sort"
	"testing"
)

func TestGoMapSetEmpty(t *testing.T) { setTest(t, set.NewGoMapSet(0), []string{}) }
func TestGoMapSetOne(t *testing.T)   { setTest(t, set.NewGoMapSet(0), []string{"A"}) }
func TestGoMapSetMany(t *testing.T)  { setTest(t, set.NewGoMapSet(0), []string{"A", "B", "C"}) }

func TestGoMapSet100Empty(t *testing.T) { setTest(t, set.NewGoMapSet(100), []string{}) }
func TestGoMapSet100One(t *testing.T)   { setTest(t, set.NewGoMapSet(100), []string{"A"}) }
func TestGoMapSet100Many(t *testing.T)  { setTest(t, set.NewGoMapSet(100), []string{"A", "B", "C"}) }

// Verifies proper implementation of a set.Set

func setTest(t *testing.T, a set.Set, want []string) {

	if !a.IsEmpty() {
		t.Fatalf("should be empty")
	}

	for i, k := range want {

		if a.Len() != i {
			t.Fatalf("should have size %d now", i)
		}

		if a.Contains(k) {
			t.Fatalf("should not contain %q just yet", k)
		}

		a.Add(k)

		if !a.Contains(k) {
			t.Fatalf("should contain %q now", k)
		}

		if a.Len() != i+1 {
			t.Fatalf("should have size %d now", i+1)
		}
	}

	if a.Len() != len(want) {
		t.Fatalf("should have size %d after insertions", len(want))
	}

	// notA := set.NewGoMapSet(setA.Len())
	// set.Difference(setA, a, notA)

	// for _, k := range notA.Keys() {
	// 	if a.Contains(k) {
	// 		t.Fatalf("should no contain %q", k)
	// 	}
	// }

	if a.Len() != len(want) {
		t.Fatalf("should have size %d checking invalid values", len(want))
	}

	// list first, mutate after (mutate changes the set)
	if listable, ok := a.(set.ListSet); ok {
		listableTest(t, listable, want)
	}

	if mutable, ok := a.(set.MutableSet); ok {
		mutableTest(t, mutable, want)
	}
}

func mutableTest(t *testing.T, a set.MutableSet, want []string) {
	// `a` contains all of `want`

	for _, k := range want {
		if !a.Contains(k) {
			t.Fatalf("should contain %q before deletion", q)
		}

		a.Delete(k)

		if a.Contains(k) {
			t.Fatalf("should NOT contain %q after deletion", q)
		}
	}
}

func listableTest(t *testing.T, a set.ListSet, want []string) {
	// `a` contains all of `want`
	got := a.Keys()

	if len(got) != len(want) {
		t.Fatalf("want len %d, got len %d", len(want), len(got))
	}

	sort.Strings(want)
	sort.Strings(got)

	for i, wantk := range want {
		gotk := got[i]
		if wantk != gotk {
			t.Errorf("index %d: want %q got %q", i, wantk, gotk)
		}
	}
}

// Assuming proper implementation of the set, this test suite covers
// all classes for which the Union/Intersection/Difference/XOR are valid.

func TestUnionEmpty(t *testing.T)      { union(t, []string{}, []string{}, []string{}) }
func TestUnionLeftEmpty(t *testing.T)  { union(t, []string{}, []string{"A"}, []string{"A"}) }
func TestUnionRightEmpty(t *testing.T) { union(t, []string{"A"}, []string{}, []string{"A"}) }
func TestUnionDisjoint(t *testing.T)   { union(t, []string{"A"}, []string{"B"}, []string{"A", "B"}) }
func TestUnionMany(t *testing.T) {
	union(t, []string{"A", "B"}, []string{"B", "C"}, []string{"A", "B", "C"})
}

func TestIntersectEmpty(t *testing.T)      { inter(t, []string{}, []string{}, []string{}) }
func TestIntersectLeftEmpty(t *testing.T)  { inter(t, []string{}, []string{"A"}, []string{}) }
func TestIntersectRightEmpty(t *testing.T) { inter(t, []string{"A"}, []string{}, []string{}) }
func TestIntersectNothing(t *testing.T)    { inter(t, []string{"A"}, []string{"B"}, []string{}) }
func TestIntersectSingle(t *testing.T) {
	inter(t, []string{"A", "B"}, []string{"B", "C"}, []string{"B"})
}

func TestDifferenceEmpty(t *testing.T)      { diff(t, []string{}, []string{}, []string{}) }
func TestDifferenceLeftEmpty(t *testing.T)  { diff(t, []string{}, []string{"A"}, []string{}) }
func TestDifferenceRightEmpty(t *testing.T) { diff(t, []string{"A"}, []string{}, []string{"A"}) }
func TestDifferenceNothing(t *testing.T)    { diff(t, []string{"A"}, []string{"B"}, []string{"A"}) }
func TestDifferenceSingle(t *testing.T) {
	diff(t, []string{"A", "B", "C"}, []string{"A", "B", "D"}, []string{"C"})
}

func TestXOREmpty(t *testing.T)      { xor(t, []string{}, []string{}, []string{}) }
func TestXORLeftEmpty(t *testing.T)  { xor(t, []string{}, []string{"A"}, []string{"A"}) }
func TestXORRightEmpty(t *testing.T) { xor(t, []string{"A"}, []string{}, []string{"A"}) }
func TestXORNothing(t *testing.T)    { xor(t, []string{"A"}, []string{"B"}, []string{"A", "B"}) }
func TestXORSingle(t *testing.T) {
	xor(t, []string{"A", "B", "C"}, []string{"A", "B", "D"}, []string{"C", "D"})
}

// relax the function definition to accept more specialized ListSet instead of Set
func relax(f func(set.ListSet, set.Set, set.Set)) func(set.ListSet, set.ListSet, set.Set) {
	return func(a, b set.ListSet, out set.Set) { f(a, b, out) }
}

func union(t *testing.T, a, b, want []string) {
	checkOp(t, a, b, want, set.Union)
}

func inter(t *testing.T, a, b, want []string) {
	checkOp(t, a, b, want, relax(set.Intersect))
}

func diff(t *testing.T, a, b, want []string) {
	checkOp(t, a, b, want, relax(set.Difference))
}

func xor(t *testing.T, a, b, want []string) {
	checkOp(t, a, b, want, set.XOR)
}

func checkOp(t *testing.T, a, b, want []string, op func(set.ListSet, set.ListSet, set.Set)) {
	A := setFromList(a)
	B := setFromList(b)
	Want := setFromList(want)
	Out := set.NewGoMapSet(len(want))

	op(A, B, Out)

	for _, k := range Want.Keys() {
		if !Out.Contains(k) {
			t.Errorf("missing %q", k)
		}
	}

	for _, k := range Out.Keys() {
		if !Want.Contains(k) {
			t.Errorf("extra %q", k)
		}
	}
}
