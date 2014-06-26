package set_test

import (
	"github.com/aybabtme/set"
	"testing"
)

func TestTchapPatricia_Collision(t *testing.T) { collisionTest(t, set.NewTchapPatricia()) }

func TestTchapPatricia_Empty(t *testing.T) { setTest(t, set.NewTchapPatricia(), []string{}) }
func TestTchapPatricia_One(t *testing.T)   { setTest(t, set.NewTchapPatricia(), []string{"A"}) }
func TestTchapPatricia_Many(t *testing.T)  { setTest(t, set.NewTchapPatricia(), []string{"A", "B", "C"}) }
func TestTchapPatricia_Operations(t *testing.T) {
	checkSetOp(t, func() set.Set { return set.NewTchapPatricia() })
}
