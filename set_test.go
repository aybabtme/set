package set_test

import (
	"github.com/aybabtme/set"
	"sort"
	"sync"
	"testing"
)

func TestGoMap_Collision(t *testing.T) { collisionTest(t, set.NewGoMap(0)) }
func TestGoMap_Empty(t *testing.T)     { setTest(t, set.NewGoMap(0), []string{}) }
func TestGoMap_One(t *testing.T)       { setTest(t, set.NewGoMap(0), []string{"A"}) }
func TestGoMap_Many(t *testing.T)      { setTest(t, set.NewGoMap(0), []string{"A", "B", "C"}) }
func TestGoMap_Operations(t *testing.T) {
	checkSetOp(t, func() set.Set { return set.NewGoMap(0) })
}

func TestGoMap_100Empty(t *testing.T) { setTest(t, set.NewGoMap(100), []string{}) }
func TestGoMap_100One(t *testing.T)   { setTest(t, set.NewGoMap(100), []string{"A"}) }
func TestGoMap_100Many(t *testing.T)  { setTest(t, set.NewGoMap(100), []string{"A", "B", "C"}) }
func TestGoMap_100Operations(t *testing.T) {
	checkSetOp(t, func() set.Set { return set.NewGoMap(100) })
}

// Verifies proper implementation of a set.Set

func collisionTest(t *testing.T, a set.Set) {
	var collisions int
	for _, word := range setA.Keys() {
		if a.Contains(word) {
			collisions++
		}
		a.Add(word)
	}
	if collisions != 0 {
		t.Errorf("%d collisions", collisions)
	}
}

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

	notA := set.NewGoMap(setA.Len())
	set.Difference(setA, a, notA)

	for _, k := range notA.Keys() {
		if a.Contains(k) {
			t.Fatalf("should no contain %q", k)
		}
	}

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
			t.Fatalf("should contain %v before deletion", k)
		}

		a.Delete(k)

		if a.Contains(k) {
			t.Fatalf("should NOT contain %v after deletion: %#v", k, a)
		}
	}
}

func listableTest(t *testing.T, a set.ListSet, want []string) {
	// `a` contains all of `want`
	got := a.Keys()

	if len(got) != len(want) {
		t.Fatalf("want %d elements from Keys(), got %d", len(want), len(got))
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
//
// The flaw in this claim is in its supposition that set.GoMapSet is correct.
// However, if set.GoMapSet is not correct, it will likely fail the part of
// it's test that don't depend on set operations.

func TestUnionEmpty(t *testing.T)      { union(t, []string{}, []string{}, []string{}) }
func TestUnionLeftEmpty(t *testing.T)  { union(t, []string{}, []string{"A"}, []string{"A"}) }
func TestUnionRightEmpty(t *testing.T) { union(t, []string{"A"}, []string{}, []string{"A"}) }
func TestUnionDisjoint(t *testing.T)   { union(t, []string{"A"}, []string{"B"}, []string{"A", "B"}) }
func TestUnionMany(t *testing.T) {
	union(t, []string{"A", "B"}, []string{"B", "C"}, []string{"A", "B", "C"})
}
func TestUnionManyDisorder(t *testing.T) {
	union(t, []string{"Z", "A", "B"}, []string{"B", "P", "Y", "C"}, []string{"A", "B", "C", "P", "Y", "Z"})
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

// If the baseline works, any Set can be used as output

type operation func(set.ListSet, set.ListSet, set.Set)

// relax the function definition to accept more specialized ListSet instead of Set
func relax(f func(set.ListSet, set.Set, set.Set)) operation {
	return func(a, b set.ListSet, out set.Set) { f(a, b, out) }
}

func union(t *testing.T, a, b, want []string) { checkOp(t, a, b, want, set.Union) }
func inter(t *testing.T, a, b, want []string) { checkOp(t, a, b, want, relax(set.Intersect)) }
func diff(t *testing.T, a, b, want []string)  { checkOp(t, a, b, want, relax(set.Difference)) }
func xor(t *testing.T, a, b, want []string)   { checkOp(t, a, b, want, set.XOR) }

func checkOp(t *testing.T, a, b, want []string, op operation) {
	checkOpBuilder(t, a, b, want, op, func() set.Set { return set.NewGoMap(0) })
}

var once *sync.Once

func checkOpBuilder(t *testing.T, a, b, want []string, op operation, outbuild func() set.Set) {
	A := setFromList(a)
	B := setFromList(b)
	Want := setFromList(want)
	Out := outbuild()

	op(A, B, Out)

	for _, k := range Want.Keys() {
		if !Out.Contains(k) {
			t.Errorf("missing %q", k)
		}
	}

	listable, ok := Out.(set.ListSet)
	if !ok {
		if once == nil {
			once = &sync.Once{}
		}
		once.Do(func() { t.Logf("weaker guarantee: %T is not listable. operations partialy tested", Out) })
		return
	}

	for _, k := range listable.Keys() {
		if !Want.Contains(k) {
			t.Errorf("extra %q", k)
		}
	}
}

// Test suite to assert that a set.Set properly supports set operations.
// This is a rehash of the the tests used to assert the operations are
// correct to begin with.

type setcase struct {
	op    operation
	cases []opcase
}
type opcase struct{ A, B, Want []string }

var setOpsTT = []setcase{
	{
		op: set.Union,
		cases: []opcase{
			{A: []string{}, B: []string{}, Want: []string{}},
			{A: []string{}, B: []string{"A"}, Want: []string{"A"}},
			{A: []string{"A"}, B: []string{}, Want: []string{"A"}},
			{A: []string{"A"}, B: []string{"B"}, Want: []string{"A", "B"}},
			{A: []string{"A", "B"}, B: []string{"B", "C"}, Want: []string{"A", "B", "C"}},
			{A: []string{"Zelda", "A", "B"}, B: []string{"B", "P", "Y", "C"}, Want: []string{"A", "B", "C", "P", "Y", "Zelda"}},
		},
	},
	{
		op: relax(set.Intersect),
		cases: []opcase{
			{A: []string{}, B: []string{}, Want: []string{}},
			{A: []string{}, B: []string{"A"}, Want: []string{}},
			{A: []string{"A"}, B: []string{}, Want: []string{}},
			{A: []string{"A"}, B: []string{"B"}, Want: []string{}},
			{A: []string{"A", "B"}, B: []string{"B", "C"}, Want: []string{"B"}},
		},
	},
	{
		op: relax(set.Difference),
		cases: []opcase{
			{A: []string{}, B: []string{}, Want: []string{}},
			{A: []string{}, B: []string{"A"}, Want: []string{}},
			{A: []string{"A"}, B: []string{}, Want: []string{"A"}},
			{A: []string{"A"}, B: []string{"B"}, Want: []string{"A"}},
			{A: []string{"A", "B", "C"}, B: []string{"A", "B", "D"}, Want: []string{"C"}},
		},
	},
	{
		op: set.XOR,
		cases: []opcase{
			{A: []string{}, B: []string{}, Want: []string{}},
			{A: []string{}, B: []string{"A"}, Want: []string{"A"}},
			{A: []string{"A"}, B: []string{}, Want: []string{"A"}},
			{A: []string{"A"}, B: []string{"B"}, Want: []string{"A", "B"}},
			{A: []string{"A", "B", "C"}, B: []string{"A", "B", "D"}, Want: []string{"C", "D"}},
		},
	},
}

func checkSetOp(t *testing.T, out func() set.Set) {
	for _, operations := range setOpsTT {
		for _, cases := range operations.cases {
			checkOpBuilder(
				t,
				cases.A,
				cases.B,
				cases.Want,
				operations.op,
				out,
			)
		}
	}
}
