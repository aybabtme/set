package set_test

import (
	"github.com/aybabtme/set"
	"sync"
	"testing"
)

var hash64table = map[string]func(int, bool) *set.Hash64{
	"Spooky64":   set.NewSpooky64,
	"Farmhash64": set.NewFarm64,
}

func TestHash64_Empty(t *testing.T) { checkHash64(t, 0, []string{}) }
func TestHash64_One(t *testing.T)   { checkHash64(t, 0, []string{"A"}) }
func TestHash64_Many(t *testing.T)  { checkHash64(t, 0, []string{"A", "B", "C"}) }
func TestHash64_Operations(t *testing.T) {
	checkHash64SetOp(t, 0)
}

func TestHash64_100Empty(t *testing.T) { checkHash64(t, 100, []string{}) }
func TestHash64_100One(t *testing.T)   { checkHash64(t, 100, []string{"A"}) }
func TestHash64_100Many(t *testing.T) {
	checkHash64(t, 100, []string{"A", "B", "C"})
}
func TestHash64_100Operations(t *testing.T) {
	checkHash64SetOp(t, 100)
}

func TestHash64_Collision(t *testing.T) { collisionHash64Test(t) }

func collisionHash64Test(t *testing.T) {
	for name, hashset := range hash64table {
		t.Logf("-- Hash64: %q --", name)
		collisionTest(t, hashset(0, true))
	}
}

func checkHash64(t *testing.T, n int, want []string) {
	for name, hashset := range hash64table {
		t.Logf("-- Hash64: %q --", name)
		once = &sync.Once{}
		setTest(t, hashset(n, true), want)
	}
}
func checkHash64SetOp(t *testing.T, n int) {
	for name, hashset := range hash64table {
		t.Logf("-- Hash64: %q --", name)
		once = &sync.Once{}
		checkSetOp(t, func() set.Set { return hashset(n, false) })
	}
}
