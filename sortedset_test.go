package set_test

import (
	"github.com/aybabtme/set"
	"testing"
)

func TestSortedSet_Empty(t *testing.T) { setTest(t, set.NewSortedSet(0), []string{}) }
func TestSortedSet_One(t *testing.T)   { setTest(t, set.NewSortedSet(0), []string{"A"}) }
func TestSortedSet_Many(t *testing.T)  { setTest(t, set.NewSortedSet(0), []string{"A", "B", "C"}) }
func TestSortedSet_Operations(t *testing.T) {
	checkSetOp(t, func() set.Set { return set.NewSortedSet(0) })
}

// unpractical for pretty much all cases
// func TestSortedSet_Collision(t *testing.T) { collisionTest(t, set.NewSortedSet(0)) }
