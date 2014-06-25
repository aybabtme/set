package set_test

import (
	"github.com/aybabtme/set"
	"sync"
	"testing"
)

var hash128table = map[string]func(int) *set.Hash128{
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

func checkHash128(t *testing.T, n int, want []string) {
	for name, hashset := range hash128table {
		t.Logf("-- Hash128: %q --", name)
		once = &sync.Once{}
		setTest(t, hashset(n), want)
	}
}
func checkHash128SetOp(t *testing.T, n int) {
	for name, hashset := range hash128table {
		t.Logf("-- Hash128: %q --", name)
		once = &sync.Once{}
		checkSetOp(t, func() set.Set { return hashset(n) })
	}
}
