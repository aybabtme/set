package set_test

import (
	"github.com/aybabtme/set"
	"testing"
)

func TestQuicktrie_Collision(t *testing.T) { collisionTest(t, set.NewQuicktrie()) }

func TestQuicktrie_Empty(t *testing.T) { setTest(t, set.NewQuicktrie(), []string{}) }
func TestQuicktrie_One(t *testing.T)   { setTest(t, set.NewQuicktrie(), []string{"A"}) }
func TestQuicktrie_Many(t *testing.T)  { setTest(t, set.NewQuicktrie(), []string{"A", "B", "C"}) }
func TestQuicktrie_Operations(t *testing.T) {
	checkSetOp(t, func() set.Set { return set.NewQuicktrie() })
}
