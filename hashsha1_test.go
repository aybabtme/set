package set_test

import (
	"github.com/aybabtme/set"
	"testing"
)

func TestHashSHA1_Collision(t *testing.T) { collisionTest(t, set.NewHashSHA1(0, true)) }

func TestHashSHA1_Empty(t *testing.T) { setTest(t, set.NewHashSHA1(0, true), []string{}) }
func TestHashSHA1_One(t *testing.T)   { setTest(t, set.NewHashSHA1(0, true), []string{"A"}) }
func TestHashSHA1_Many(t *testing.T)  { setTest(t, set.NewHashSHA1(0, true), []string{"A", "B", "C"}) }
func TestHashSHA1_Operations(t *testing.T) {
	checkSetOp(t, func() set.Set { return set.NewHashSHA1(0, false) })
}

func TestHashSHA1_100Empty(t *testing.T) { setTest(t, set.NewHashSHA1(100, true), []string{}) }
func TestHashSHA1_100One(t *testing.T)   { setTest(t, set.NewHashSHA1(100, true), []string{"A"}) }
func TestHashSHA1_100Many(t *testing.T) {
	setTest(t, set.NewHashSHA1(100, true), []string{"A", "B", "C"})
}
func TestHashSHA1_100Operations(t *testing.T) {
	checkSetOp(t, func() set.Set { return set.NewHashSHA1(100, false) })
}
