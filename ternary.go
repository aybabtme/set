package set

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type ternNode struct {
	id     int
	exists bool
	Code   rune

	left  *ternNode
	child *ternNode
	right *ternNode
}

// TernarySet is a set specifically for string indexed keys.
type TernarySet struct {
	root  *ternNode
	count int
	ids   int
}

var (
	ternarySetIsListable ListSet = NewTernarySet()
)

// NewTernarySet creates a trie.
func NewTernarySet() *TernarySet {
	return &TernarySet{nil, 0, 0}
}

// Add the key to the set.
func (t *TernarySet) Add(key string) {

	t.root = t.put(t.root, key, 0)
}

func (t *TernarySet) put(x *ternNode, key string, d int) *ternNode {
	c := key[d]
	if x == nil {
		x = &ternNode{id: t.ids, Code: rune(c)}
		t.ids++
	}

	if c < uint8(x.Code) {
		x.left = t.put(x.left, key, d)
	} else if c > uint8(x.Code) {
		x.right = t.put(x.right, key, d)
	} else if d < len(key)-1 {
		x.child = t.put(x.child, key, d+1)
	} else {
		x.exists = true
		t.count++
	}
	return x
}

// Contains tells if key exists.
func (t *TernarySet) Contains(key string) bool {

	if key == "" {
		return false
	}

	ternNode := t.get(t.root, key, 0)

	if ternNode == nil {
		return false
	}

	return ternNode.exists
}

func (t *TernarySet) get(x *ternNode, key string, d int) *ternNode {
	if x == nil {
		return nil
	}

	c := key[d]

	if c < uint8(x.Code) {
		return t.get(x.left, key, d)
	} else if c > uint8(x.Code) {
		return t.get(x.right, key, d)
	} else if d < len(key)-1 {
		return t.get(x.child, key, d+1)
	}

	return x
}

// Len returns the count of elements in this set.
func (t *TernarySet) Len() int { return t.count }

// IsEmpty tells if this set contains any elements.
func (t *TernarySet) IsEmpty() bool { return t.Len() == 0 }

// Keys returns all the keys known to this trie
func (t *TernarySet) Keys() []string {
	var outCollection []string

	t.Debug(os.Stderr, "hello")
	collect(t.root, []uint8{}, outCollection)

	return outCollection
}

func (t *TernarySet) Debug(out io.Writer, name string) {
	buf := bytes.NewBuffer(nil)

	_, _ = fmt.Fprintf(out, "digraph %s {\n", name)
	visit(t.root, out, buf)
	_, _ = buf.WriteTo(out)
	_, _ = fmt.Fprintf(out, "}\n")
}

func visit(x *ternNode, nodes, edges io.Writer) {

	if x == nil {
		_, _ = fmt.Fprintln(edges, "nil;")
		return
	} else {
		_, _ = fmt.Fprintf(edges, "%d;\n", x.id)
	}

	if x.exists {
		fmt.Fprintf(nodes, "\t%d [label=\"%c\", shape = doublecircle];\n", x.id, x.Code)
	} else {
		fmt.Fprintf(nodes, "\t%d [label=\"%c\", shape = circle];\n", x.id, x.Code)
	}

	_, _ = fmt.Fprintf(edges, "\t%d -> ", x.id)
	visit(x.left, nodes, edges)
	_, _ = fmt.Fprintf(edges, "\t%d -> ", x.id)
	visit(x.child, nodes, edges)
	_, _ = fmt.Fprintf(edges, "\t%d -> ", x.id)
	visit(x.right, nodes, edges)

}

// Helpers

func collect(x *ternNode, key []byte, outCollection []string) {
	if x == nil {
		return
	}
	collect(x.left, key, outCollection)
	newKey := append(key, uint8(x.Code))
	if x.exists {
		log.Printf("collect: %s", string(newKey))
		outCollection = append(outCollection, string(newKey))
	} else {
		log.Printf("NOT collect: %s", string(newKey))
	}
	collect(x.child, newKey, outCollection)
	collect(x.right, key, outCollection)
}
