package set_test

import (
	"bufio"
	"github.com/aybabtme/set"
	"log"
	"os"
)

var (
	q     = struct{}{}
	web2  = loadWordlist("/usr/share/dict/web2", web2fallback)
	web2a = loadWordlist("/usr/share/dict/web2a", web2afallback)

	setA = setFromList(web2)
	setB = setFromList(web2a)
)

func init() {
	A := setFromList(web2)
	B := setFromList(web2a)

	log.Printf("|A| = %d", A.Len())
	log.Printf("|B| = %d", B.Len())

	u := set.NewGoMap(0)
	set.Union(A, B, u)
	log.Printf("|A ∪ B| = %d", u.Len())

	i := set.NewGoMap(0)
	set.Intersect(A, B, i)
	log.Printf("|A ∩ B| = %d", i.Len())

	d := set.NewGoMap(0)
	set.Difference(A, B, d)
	log.Printf("|A - B| = %d", d.Len())

	x := set.NewGoMap(0)
	set.XOR(A, B, x)
	log.Printf("|A ⊕ B| = %d", x.Len())
}

func loadWordlist(filename string, fallback []string) (out []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("couldn't open %q: %v", filename, err)
		return fallback
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		out = append(out, scan.Text())
	}

	if scan.Err() != nil {
		log.Printf("scanning %q: %v", filename, scan.Err())
		return fallback
	}

	log.Printf("%q: %d words", filename, len(out))

	return
}

func setFromList(words []string) set.GoMap {
	s := set.NewGoMap(len(words))
	for _, word := range words {
		s[word] = q
	}
	return s
}

var (
	web2fallback  = []string{"A", "a", "aa", "aal", "aalii", "aam", "Aani", "aardvark", "aardwolf", "Aaron", "Aaronic", "Aaronical", "Aaronite", "Aaronitic", "Aaru", "Ab", "aba", "Ababdeh", "Ababua", "abac", "abaca", "abacate", "abacay", "abacinate", "abacination", "abaciscus", "abacist", "aback", "abactinal", "abactinally", "abaction", "abactor", "abaculus", "abacus", "Abadite", "abaff", "abaft", "abaisance", "abaiser", "abaissed", "abalienate", "abalienation", "abalone", "Abama", "abampere", "abandon", "abandonable", "abandoned", "abandonedly", "abandonee", "abandoner", "abandonment", "Abanic", "Abantes", "abaptiston", "Abarambo", "Abaris", "abarthrosis", "abarticular", "abarticulation", "abas", "abase", "abased", "abasedly", "abasedness", "abasement", "abaser", "Abasgi", "abash", "abashed", "abashedly", "abashedness", "abashless", "abashlessly", "abashment", "abasia", "abasic", "abask", "Abassin", "abastardize", "abatable", "abate", "abatement", "abater", "abatis", "abatised", "abaton", "abator", "abattoir", "Abatua", "abature", "abave", "abaxial", "abaxile", "abaze", "abb", "Abba", "abbacomes", "abbacy", "Abbadide"}
	web2afallback = []string{"A", "acid", "abacus", "major", "abacus", "pythagoricus", "A", "battery", "abbey", "counter", "abbey", "laird", "abbey", "lands", "abbey", "lubber", "abbot", "cloth", "Abbott", "papyrus", "abb", "wool", "A-b-c", "book", "A-b-c", "method", "abdomino-uterotomy", "Abdul-baha", "a-be", "aberrant", "duct", "aberration", "constant", "abiding", "place", "able-bodied", "able-bodiedness", "able-minded", "able-mindedness", "able", "seaman", "aboli", "fruit", "A", "bond", "Abor-miri", "a-borning", "about-face", "about", "ship", "about-sledge", "above-cited", "above-found", "above-given", "above-mentioned", "above-named", "above-quoted", "above-reported", "above-said", "above-water", "above-written", "Abraham-man", "abraum", "salts", "abraxas", "stone", "Abri", "audit", "culture", "abruptly", "acuminate", "abruptly", "pinnate", "absciss", "layer", "absence", "state", "absentee", "voting", "absent-minded", "absent-mindedly", "absent-mindedness", "absent", "treatment", "absent", "voter", "Absent", "voting", "absinthe", "green", "absinthe", "oil", "absorption", "bands", "absorption", "circuit", "absorption", "coefficient", "absorption", "current", "absorption", "dynamometer", "absorption", "factor", "absorption", "lines", "absorption", "pipette", "absorption", "screen", "absorption", "spectrum", "absorption", "system", "A", "b", "station", "abstinence", "theory", "abstract", "group", "Abt", "system", "abundance", "declaree", "aburachan", "seed", "abutment", "arch", "abutment", "pier", "abutting", "joint", "acacia", "veld", "academy", "blue", "academy", "board", "academy", "figure", "acajou", "balsam", "acanthosis", "nigricans", "acanthus", "family", "acanthus", "leaf", "acaroid", "resin", "Acca", "larentia", "acceleration", "note", "accelerator", "nerve", "accent", "mark", "acceptance", "bill", "acceptance", "house", "acceptance", "supra", "protest", "acceptor", "supra", "protest", "accession", "book", "accession", "number", "accession", "service", "access", "road", "accident", "insurance"}
)
