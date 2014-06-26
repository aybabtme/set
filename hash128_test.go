package set_test

import (
	"github.com/aybabtme/set"
	"sync"
	"testing"
)

var hash128table = map[string]func(int, bool) *set.Hash128{
	"Spooky128":   set.NewSpooky128,
	"Farmhash128": set.NewFarm128,
}

func TestHash128_Empty(t *testing.T) { checkHash128(t, 0, []string{}) }
func TestHash128_One(t *testing.T)   { checkHash128(t, 0, []string{"A"}) }
func TestHash128_Many(t *testing.T)  { checkHash128(t, 0, []string{"A", "B", "C"}) }
func TestHash128_Operations(t *testing.T) {
	checkHash128SetOp(t, 0)
}

func TestHash128_100Empty(t *testing.T) { checkHash128(t, 100, []string{}) }
func TestHash128_100One(t *testing.T)   { checkHash128(t, 100, []string{"A"}) }
func TestHash128_100Many(t *testing.T) {
	checkHash128(t, 100, []string{"A", "B", "C"})
}
func TestHash128_100Operations(t *testing.T) {
	checkHash128SetOp(t, 100)
}

func TestHash128_Collision(t *testing.T) { collisionHash128Test(t) }

func collisionHash128Test(t *testing.T) {
	for name, hashset := range hash128table {
		t.Logf("-- Hash128: %q --", name)
		collisionTest(t, hashset(0, true))
	}
}

func checkHash128(t *testing.T, n int, want []string) {
	for name, hashset := range hash128table {
		t.Logf("-- Hash128: %q --", name)
		once = &sync.Once{}
		setTest(t, hashset(n, true), want)
	}
}
func checkHash128SetOp(t *testing.T, n int) {
	for name, hashset := range hash128table {
		t.Logf("-- Hash128: %q --", name)
		once = &sync.Once{}
		checkSetOp(t, func() set.Set { return hashset(n, false) })
	}
}
