package main

import (
	"fmt"
	"github.com/aybabtme/benchkit"
	"github.com/aybabtme/benchkit/benchplot"
	"github.com/aybabtme/set"
	"github.com/aybabtme/uniplot/spark"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func timeplotCommand() ([]cli.Flag, func(*cli.Context)) {

	fileFlag := cli.StringFlag{Name: "file", Usage: "file containing the keys to read from"}
	typeFlag := cli.StringFlag{Name: "type", Usage: "type of set to benchmark"}
	widthFlag := cli.Float64Flag{Name: "width", Value: 16.0, Usage: "width of the plot to render"}
	heightFlag := cli.Float64Flag{Name: "height", Value: 12.0, Usage: "height of the plot to render"}

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

		log.Printf("key-count=%d", len(keys))

		log.Printf("doing %d benchmarks", len(sets))
		for i, s := range sets {
			runtime.GC()
			doContainsEmptyBenchmark(s.name, s.s, keys, filename, width, height)
			doAddBenchmark(s.name, s.s, keys, filename, width, height)
			doContainsFullBenchmark(s.name, s.s, keys, filename, width, height)
			log.Printf("%d/%d done", i+1, len(sets))
		}

	}
}

func doContainsEmptyBenchmark(setName string, setter func() set.Set, keys []string, filename string, width, height float64) {
	keyfactor := 100

	start := time.Now()
	n := len(keys)/keyfactor + 2
	m := 30

	log.Printf("benchmarking Contains on empty %q...", setName)

	results := benchkit.Bench(benchkit.Time(n, m)).Each(func(each benchkit.BenchEach) {
		for repeat := 0; repeat < m && !abort; repeat++ {
			log.Printf("doing pass %d/%d", repeat+1, m)
			lastj := 0
			set := setter()
			runtime.GC()
			each.Before(0)
			for j, key := range keys {
				if abort {
					return
				}

				if j%keyfactor == 0 {
					each.After(lastj)
					each.Before(lastj + 1)
					lastj++
				}
				set.Contains(key)

			}
			each.After(lastj)
		}
	}).(*benchkit.TimeResult)
	log.Printf("done in %v", time.Since(start))

	plottitle := fmt.Sprintf("Membership of %d keys in an empty %s", len(keys), setName)
	log.Printf("plotting %q", plottitle)
	p, err := benchplot.PlotTime(plottitle, fmt.Sprintf("times %d keys", keyfactor), results, true)
	if err != nil {
		log.Printf("plotting=%q\terror=%v", plottitle, err)
		return
	}

	base := filepath.Base(filename)
	ext := filepath.Ext(base)
	cleanname := base[:len(base)-len(ext)]

	plotname := fmt.Sprintf("timeplot_%s_contains_empty_%s", cleanname, setName)
	log.Printf("saving to %q (%gx%g)", plotname, width, height)

	if err := p.Save(width, height, plotname+".png"); err != nil {
		log.Printf("saving=%q\terror=%v", plotname, err)
	}
}

func doAddBenchmark(setName string, setter func() set.Set, keys []string, filename string, width, height float64) {
	keyfactor := 100

	start := time.Now()
	n := len(keys)/keyfactor + 2
	m := 30

	log.Printf("benchmarking Add on %q...", setName)

	results := benchkit.Bench(benchkit.Time(n, m)).Each(func(each benchkit.BenchEach) {
		for repeat := 0; repeat < m && !abort; repeat++ {
			log.Printf("doing pass %d/%d", repeat+1, m)
			lastj := 0
			set := setter()
			runtime.GC()
			each.Before(0)
			for j, key := range keys {
				if abort {
					return
				}
				if j%keyfactor == 0 {
					each.After(lastj)
					each.Before(lastj + 1)
					lastj++
				}
				set.Add(key)

			}
			each.After(lastj)
		}
	}).(*benchkit.TimeResult)
	log.Printf("done in %v", time.Since(start))

	plottitle := fmt.Sprintf("Insertion of %d keys in a %s", len(keys), setName)
	log.Printf("plotting %q", plottitle)
	p, err := benchplot.PlotTime(plottitle, fmt.Sprintf("times %d keys", keyfactor), results, true)
	if err != nil {
		log.Printf("plotting=%q\terror=%v", plottitle, err)
		return
	}

	base := filepath.Base(filename)
	ext := filepath.Ext(base)
	cleanname := base[:len(base)-len(ext)]

	plotname := fmt.Sprintf("timeplot_%s_add_%s", cleanname, setName)
	log.Printf("saving to %q (%gx%g)", plotname, width, height)

	if err := p.Save(width, height, plotname+".png"); err != nil {
		log.Printf("saving=%q\terror=%v", plotname, err)
	}
}

func doContainsFullBenchmark(setName string, setter func() set.Set, keys []string, filename string, width, height float64) {
	keyfactor := 100

	start := time.Now()
	n := len(keys)/keyfactor + 2
	m := 30

	log.Printf("benchmarking Contains on full %q...", setName)

	results := benchkit.Bench(benchkit.Time(n, m)).Each(func(each benchkit.BenchEach) {
		for repeat := 0; repeat < m && !abort; repeat++ {
			log.Printf("doing pass %d/%d", repeat+1, m)
			lastj := 0
			set := setter()

			// fill the set
			for _, key := range keys {
				if abort {
					return
				}
				set.Add(key)
			}

			runtime.GC()
			each.Before(0)
			for j, key := range keys {
				if abort {
					return
				}
				if j%keyfactor == 0 {
					each.After(lastj)
					each.Before(lastj + 1)
					lastj++
				}
				set.Contains(key)

			}
			each.After(lastj)
		}
	}).(*benchkit.TimeResult)
	log.Printf("done in %v", time.Since(start))

	plottitle := fmt.Sprintf("Membership of %d keys in an full %s", len(keys), setName)
	log.Printf("plotting %q", plottitle)
	p, err := benchplot.PlotTime(plottitle, fmt.Sprintf("times %d keys", keyfactor), results, true)
	if err != nil {
		log.Printf("plotting=%q\terror=%v", plottitle, err)
		return
	}

	base := filepath.Base(filename)
	ext := filepath.Ext(base)
	cleanname := base[:len(base)-len(ext)]

	plotname := fmt.Sprintf("timeplot_%s_contains_full_%s", cleanname, setName)
	log.Printf("saving to %q (%gx%g)", plotname, width, height)

	if err := p.Save(width, height, plotname+".png"); err != nil {
		log.Printf("saving=%q\terror=%v", plotname, err)
	}
}
