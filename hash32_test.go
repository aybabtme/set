package set_test

import (
	"github.com/aybabtme/set"
	"sync"
	"testing"
)

var hash32table = map[string]func(int) *set.Hash32{
	"Spooky32":      set.NewSpooky32,
	"Farmhash32":    set.NewFarm32,
	"Adler32":       set.NewAdler32,
	"Murmur32":      set.NewMurmur32,
	"Djb32":         set.NewDjb32,
	"Elf32":         set.NewElf32,
	"Java32":        set.NewJava32,
	"Jenkins32":     set.NewJenkins32,
	"SDBM32":        set.NewSDBM32,
	"SQLite32":      set.NewSQLite32,
	"SuperFastHash": set.NewSuperFastHash,
}

func TestHash32_Empty(t *testing.T) { checkHash32(t, 0, []string{}) }
func TestHash32_One(t *testing.T)   { checkHash32(t, 0, []string{"A"}) }
func TestHash32_Many(t *testing.T)  { checkHash32(t, 0, []string{"A", "B", "C"}) }
func TestHash32_Operations(t *testing.T) {
	checkHash32Op(t, 0)
}

func TestHash32_100Empty(t *testing.T) { checkHash32(t, 100, []string{}) }
func TestHash32_100One(t *testing.T)   { checkHash32(t, 100, []string{"A"}) }
func TestHash32_100Many(t *testing.T) {
	checkHash32(t, 100, []string{"A", "B", "C"})
}
func TestHash32_100Operations(t *testing.T) {
	checkHash32Op(t, 100)
}

func checkHash32(t *testing.T, n int, want []string) {
	for name, hashset := range hash32table {
		t.Logf("-- Hash32: %q --", name)
		once = &sync.Once{}
		setTest(t, hashset(n), want)
	}
}
func checkHash32Op(t *testing.T, n int) {
	for name, hashset := range hash32table {
		t.Logf("-- Hash32: %q --", name)
		once = &sync.Once{}
		checkSetOp(t, func() set.Set { return hashset(n) })
	}
}
