package main

import (
	"bufio"
	"bytes"
	"compress/flate"
	"compress/gzip"
	"compress/lzw"
	"compress/zlib"
	"flag"
	"fmt"
	"github.com/aybabtme/uniplot/histogram"
	"github.com/aybabtme/uniplot/spark"
	"github.com/cheggaaa/pb"
	"github.com/dustin/go-humanize"
	"io"
	"log"
	"math"
	"os"
	"runtime"
	"sort"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetPrefix("[encbench] ")
	log.SetFlags(0)
	var (
		encoding = flag.String("enc", "lzw", "encoding algorithm to benchmark")
	)
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatalf("need a unique argument (filename)")
	}

	filename := flag.Args()[0]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("opening %q: %v", filename, err)
	}
	defer file.Close()

	scan := bufio.NewScanner(spark.Reader(file))
	scan.Split(bufio.ScanLines)

	log.Printf("reading lines from %q", filename)
	var lines [][]byte
	for scan.Scan() {
		lines = append(lines, scan.Bytes())
	}
	if err := scan.Err(); err != nil {
		log.Fatalf("scanning lines in %q: %v", filename, err)
	}

	log.Printf("compressing %d lines...", len(lines))
	report, err := compressLines(lines, *encoding)
	if err != nil {
		log.Fatalf("compressing lines: %v", err)
	}
	printStats(report)

}

func compressLines(lines [][]byte, encoding string) ([]Report, error) {
	var wc io.WriteCloser
	var start time.Time
	buf := bytes.NewBuffer(make([]byte, 0, 8096))

	reports := make([]Report, len(lines))

	bar := pb.New(len(lines))
	bar.SetUnits(pb.U_NO)
	bar.ShowTimeLeft = true
	bar.ShowSpeed = true
	bar.Start()

	var before, after uint64
	for i, line := range lines {
		switch encoding {
		case "lzw":
			wc = lzw.NewWriter(buf, lzw.LSB, 8)
		case "gzip":
			wc = gzip.NewWriter(buf)
		case "flate":
			w, err := flate.NewWriter(buf, flate.DefaultCompression)
			if err != nil {
				return nil, err
			}
			wc = w
		case "zlib":
			wc = zlib.NewWriter(buf)
		default:
			return nil, fmt.Errorf("bad encoding type: %q", encoding)
		}

		start = time.Now()
		_, err := wc.Write(line)
		if err != nil {
			log.Fatalf("failed to compress line %d: %v", i, err)
		}
		if err := wc.Close(); err != nil {
			log.Fatalf("failed to close compressed line %d: %v", i, err)
		}
		reports[i].dT = time.Since(start)
		reports[i].from = len(line)
		reports[i].to = buf.Len()
		reports[i].data = buf.Bytes()
		before += uint64(reports[i].from)
		after += uint64(reports[i].to)
		buf.Reset()
		bar.Increment()
	}
	bar.Finish()

	log.Printf("before=%s\t after=%s", humanize.Bytes(before), humanize.Bytes(after))

	return reports, nil
}

func printStats(reps []Report) {

	byRatio := reportByRatio(reps)
	sort.Sort(byRatio)

	log.Printf("By precompression size")
	err := histogram.Fprintf(
		os.Stdout,
		histogram.Hist(10, byRatio.Froms()),
		histogram.Linear(40),
		func(v float64) string {
			return humanize.Bytes(uint64(v))
		})
	if err != nil {
		log.Fatalf("plotting by ratio: %v", err)
	}

	log.Printf("By compression size")
	err = histogram.Fprintf(
		os.Stdout,
		histogram.Hist(10, Significant(byRatio.Tos())),
		histogram.Linear(40),
		func(v float64) string {
			return humanize.Bytes(uint64(v))
		})
	if err != nil {
		log.Fatalf("plotting by ratio: %v", err)
	}

	log.Printf("By compression ratio")
	err = histogram.Fprintf(
		os.Stdout,
		histogram.Hist(10, Significant(byRatio.Ratios())),
		histogram.Linear(40),
		func(v float64) string {
			return fmt.Sprintf("%.1g%%", v*100.0)
		})
	if err != nil {
		log.Fatalf("plotting by ratio: %v", err)
	}

	log.Printf("By shaved off bytes")
	err = histogram.Fprintf(
		os.Stdout,
		histogram.Hist(10, byRatio.SavedBytes()),
		histogram.Linear(40),
		func(v float64) string {
			return humanize.Bytes(uint64(v))
		})
	if err != nil {
		log.Fatalf("plotting by ratio: %v", err)
	}

	log.Printf("By compression time")
	byTime := reportByTime(reps)
	sort.Sort(&byTime)
	err = histogram.Fprintf(
		os.Stdout,
		histogram.Hist(10, Significant(byTime.Times())),
		histogram.Linear(40),
		func(v float64) string {
			return time.Duration(v).String()
		})
	if err != nil {
		log.Fatalf("plotting by ratio: %v", err)
	}
}

type Report struct {
	data []byte
	from int
	to   int
	dT   time.Duration
}

func (r *Report) Ratio() float64 {
	return float64(r.from-r.to) / float64(r.from)
}

type reportByRatio []Report

func (r reportByRatio) Ratios() []float64 {
	ratios := make([]float64, len(r))
	for i, rep := range r {
		ratios[i] = rep.Ratio()
	}
	return ratios
}

func (r reportByRatio) Froms() []float64 {
	ratios := make([]float64, len(r))
	for i, rep := range r {
		ratios[i] = float64(rep.from)
	}
	return ratios
}

func (r reportByRatio) Tos() []float64 {
	ratios := make([]float64, len(r))
	for i, rep := range r {
		ratios[i] = float64(rep.to)
	}
	return ratios
}

func (r reportByRatio) SavedBytes() []float64 {
	ratios := make([]float64, len(r))
	for i, rep := range r {
		ratios[i] = float64(rep.from - rep.to)
	}
	return ratios
}

func (r reportByRatio) Less(i, j int) bool { return r[i].Ratio() < r[j].Ratio() }
func (r reportByRatio) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r reportByRatio) Len() int           { return len(r) }

type reportByTime []Report

func (r reportByTime) Times() []float64 {
	ratios := make([]float64, len(r))
	for i, rep := range r {
		ratios[i] = float64(rep.dT)
	}
	return ratios
}

func (r reportByTime) Less(i, j int) bool { return r[i].dT < r[j].dT }
func (r reportByTime) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r reportByTime) Len() int           { return len(r) }

// µ is the expected value. Greek letters because we can.
func µ(all []float64) float64 {
	// since all values are equaly probable, µ is sum/length
	var sum float64
	for _, dur := range all {
		sum += dur
	}
	return sum / float64(len(all))
}

// σ is the standard deviation. Greek letters because we can.
func σ(all []float64) float64 {
	var sum float64
	µ := µ(all)
	for _, dur := range all {
		sum += ((dur - µ) * (dur - µ))
	}
	scaled := sum / float64(len(all))

	σ := math.Sqrt(scaled)

	return float64(σ)
}

// P returns the percentile duration of the step, such as p50, p90, p99,
// if all is sorted. If all is not sorted, you get garbage.
func P(all []float64, factor float64) float64 {
	if len(all) == 0 {
		return 0
	}
	pIdx := pIndex(len(all), factor)
	return all[pIdx-1]
}

// Significant returns the slice from -3σ to +3σ from the mean
// if all is sorted. If all is not sorted, you get garbage.
func Significant(all []float64) []float64 {
	if len(all) == 0 {
		return all
	}
	minIdx := pIndex(len(all), 1-0.9973)
	maxIdx := pIndex(len(all), 0.9973)
	return all[minIdx-1 : maxIdx]
}

func pIndex(base int, factor float64) int {
	power := math.Log10(factor)
	closest := 10 * math.Pow10(int(power))
	idx := int(math.Ceil((factor * float64(base)) / closest))
	return idx
}
