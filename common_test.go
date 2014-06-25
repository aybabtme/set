package set_test

import (
	"bufio"
	"github.com/aybabtme/set"
	"log"
	"os"
)

var (
	q     = struct{}{}
	web2  = loadWordlist("/usr/share/dict/web2")
	web2a = loadWordlist("/usr/share/dict/web2a")

	setA = setFromList(web2)
	setB = setFromList(web2a)
)

func init() {
	A := setFromList(web2)
	B := setFromList(web2a)

	log.Printf("|A| = %d", A.Len())
	log.Printf("|B| = %d", B.Len())

	u := set.NewGoMapSet(0)
	set.Union(A, B, u)
	log.Printf("|A ∪ B| = %d", u.Len())

	i := set.NewGoMapSet(0)
	set.Intersect(A, B, i)
	log.Printf("|A ∩ B| = %d", i.Len())

	d := set.NewGoMapSet(0)
	set.Difference(A, B, d)
	log.Printf("|A - B| = %d", d.Len())

	x := set.NewGoMapSet(0)
	set.XOR(A, B, x)
	log.Printf("|A ⊕ B| = %d", x.Len())
}

func loadWordlist(filename string) (out []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("couldn't open %q: %v", filename, err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		out = append(out, scan.Text())
	}

	if scan.Err() != nil {
		log.Fatalf("scanning %q: %v", filename, scan.Err())
	}

	log.Printf("%q: %d words", filename, len(out))

	return
}

func setFromList(words []string) set.GoMapSet {
	s := set.NewGoMapSet(len(words))
	for _, word := range words {
		s[word] = q
	}
	return s
}
