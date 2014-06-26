package set_test

import (
	"github.com/aybabtme/set"
	"testing"
)

func TestTernary_Collision(t *testing.T) { collisionTest(t, set.NewTernarySet()) }

func TestTernary_Empty(t *testing.T) { setTest(t, set.NewTernarySet(), []string{}) }
func TestTernary_One(t *testing.T)   { setTest(t, set.NewTernarySet(), []string{"A"}) }
func TestTernary_Many(t *testing.T)  { setTest(t, set.NewTernarySet(), []string{"A", "B", "C"}) }
func TestTernary_Operations(t *testing.T) {
	checkSetOp(t, func() set.Set { return set.NewTernarySet() })
}
