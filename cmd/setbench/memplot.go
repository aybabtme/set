package main

import (
	"fmt"
	"github.com/aybabtme/benchkit"
	"github.com/aybabtme/benchkit/benchplot"
	"github.com/aybabtme/set"
	"github.com/aybabtme/uniplot/spark"
	"github.com/codegangsta/cli"
	"github.com/dustin/go-humanize"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func memplotCommand() ([]cli.Flag, func(*cli.Context)) {

	fileFlag := cli.StringFlag{Name: "file", Usage: "file containing the keys to read from"}
	typeFlag := cli.StringFlag{Name: "type", Usage: "type of set to benchmark"}
	widthFlag := cli.Float64Flag{Name: "width", Value: 8.0, Usage: "width of the plot to render"}
	heightFlag := cli.Float64Flag{Name: "height", Value: 6.0, Usage: "height of the plot to render"}

	flags := []cli.Flag{fileFlag, typeFlag, widthFlag, heightFlag}

	return flags, func(c *cli.Context) {
		var (
			settype  = c.String(typeFlag.Name)
			filename = c.String(fileFlag.Name)
			width    = c.Float64(widthFlag.Name)
			height   = c.Float64(heightFlag.Name)
		)

		hadError := true
		switch {
		case settype == "":
			log.Println("Missing value for", typeFlag.Name)
		case filename == "":
			log.Println("Missing value for", fileFlag.Name)
		default:
			hadError = false
		}
		if hadError {
			cli.ShowCommandHelp(c, c.Command.Name)
			return
		}

		sets, err := setImpl(settype)
		if err != nil {
			log.Printf("Bad type flag: %v", err)
			return
		}

		file, err := os.Open(filename)
		if err != nil {
			log.Printf("opening=%q\terror=%v", filename, err)
			return
		}
		defer func() { _ = file.Close() }()

		keys, err := decodeKeys(spark.Reader(file))
		if err != nil {
			log.Printf("decoding=%q\terror=%v", filename, err)
			return
		}
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		log.Printf("baseline-in-use=%s", humanize.Bytes(mem.HeapAlloc-mem.HeapReleased))

		log.Printf("key-count=%d", len(keys))

		log.Printf("doing %d benchmarks", len(sets))
		for i, s := range sets {
			runtime.GC()
			doMemBenchmark(s.name, s.s, keys, filename, width, height)
			log.Printf("%d/%d done", i+1, len(sets))
		}

	}
}

func doMemBenchmark(setName string, setter func() set.Set, keys []string, filename string, width, height float64) {
	keyfactor := 100

	start := time.Now()
	n := len(keys)/keyfactor + 2

	log.Printf("benchmarking %q...", setName)

	results := benchkit.Bench(benchkit.Memory(n)).Each(func(each benchkit.BenchEach) {
		lastj := 0
		set := setter()
		runtime.GC()
		each.Before(0)
		for j, key := range keys {

			if j%keyfactor == 0 {
				each.After(lastj)
				each.Before(lastj + 1)
				lastj++
			}
			set.Add(key)

		}
		each.After(lastj)
	}).(*benchkit.MemResult)
	log.Printf("done in %v", time.Since(start))

	plottitle := fmt.Sprintf("Insertion of %d keys in a %s", len(keys), setName)
	log.Printf("plotting %q", plottitle)
	p, err := benchplot.PlotMemory(plottitle, fmt.Sprintf("times %d keys", keyfactor), results, false)
	if err != nil {
		log.Printf("plotting=%q\terror=%v", plottitle, err)
		return
	}

	base := filepath.Base(filename)
	ext := filepath.Ext(base)
	cleanname := base[:len(base)-len(ext)]

	plotname := fmt.Sprintf("%s_%s.svg", cleanname, setName)
	log.Printf("saving to %q (%gx%g)", plotname, width, height)
	if err := p.Save(width, height, plotname); err != nil {
		log.Printf("saving=%q\terror=%v", plotname, err)
	}
}
