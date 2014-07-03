package main

import (
	"bufio"
	"bytes"
	"flag"
	"github.com/aybabtme/uniplot/histogram"
	"github.com/dustin/go-humanize"
	"io"
	"log"
	"os"
	"sort"
)

func main() {
	nP := flag.Int("n", 1, "length of words to count")
	topP := flag.Int("top", 50, "top words occuring most frequently")
	ignoreP := flag.String("ignore", "", "ignore a word while counting")
	flag.Parse()
	// N = NP lol
	n := *nP
	top := *topP
	ignore := []byte(*ignoreP)
	shouldIgnore := len(ignore) != 0

	rd := bufio.NewReader(os.Stdin)
	count := make(map[string]int64)
	word := make([]byte, n)
	var lineLengths []float64
reading:
	for {
		line, err := rd.ReadBytes('\n')
		switch err {
		case io.EOF:
			break reading
		case nil:
			// carry on
		default:
			log.Fatal(err)
		}
		lineLengths = append(lineLengths, float64(len(line)))

		for i := 0; i < len(line)-n; i++ {
			word = line[i : i+n]
			if !shouldIgnore {
				count[string(word)]++
			} else if !bytes.Contains(word, ignore) {
				count[string(word)]++
			}

		}
	}

	var totalWord int64
	wl := make(WordList, 0, len(count))
	for k, v := range count {
		wl = append(wl, WordCount{
			word:  k,
			count: v,
		})
		totalWord += v
	}

	sort.Sort(wl)

	var topkWord int64
	for _, wc := range wl[imax(len(wl)-top, 0):] {
		log.Printf("\t%q\t%s", wc.word, humanize.Comma(wc.count))
		topkWord += wc.count
	}

	if shouldIgnore {
		log.Printf("ignoring %q", string(ignore))
	}

	log.Printf("counts: %d unique strings of length %d, top %d are %.1f%% of that.",
		len(count),
		n,
		top,
		100*float64(topkWord)/float64(totalWord),
	)

	h := histogram.Hist(top, lineLengths)
	histogram.Fprintf(os.Stdout, h, histogram.Linear(70), func(v float64) string {
		return humanize.Comma(int64(v))
	})
}

func imax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type WordCount struct {
	word  string
	count int64
}

type WordList []WordCount

func (w WordList) Less(i, j int) bool { return w[i].count < w[j].count }
func (w WordList) Len() int           { return len(w) }
func (w WordList) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
