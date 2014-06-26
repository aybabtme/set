package main

import (
	"bufio"
	"fmt"
	"github.com/aybabtme/set"
	"github.com/codegangsta/cli"
	"io"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strings"
)

var (
	abort bool
)

func NewApp() *cli.App {

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		<-c
		abort = true
	}()

	app := cli.NewApp()
	app.Name = "setbench"
	app.Usage = "Benchmarks different properties of set implementations."

	memplotFlags, memplotAction := memplotCommand()
	timeplotFlags, timeplotAction := timeplotCommand()

	app.Commands = []cli.Command{
		{
			Name:   "memplot",
			Usage:  "Plots memory usage over time as keys are inserted in a set.",
			Flags:  memplotFlags,
			Action: memplotAction,
		},
		{
			Name:   "timeplot",
			Usage:  "Plots time per operation as keys are inserted in a set.",
			Flags:  timeplotFlags,
			Action: timeplotAction,
		},
	}

	return app
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetPrefix("[setbench] ")
	log.SetFlags(log.Lshortfile)

	err := NewApp().Run(os.Args)
	if err != nil {
		log.Fatalf("Running app: %v", err)
	}
}

type setimpl struct {
	name string
	s    func() set.Set
}

var impls = map[string]setimpl{
	"gomap":       {name: "GoMap", s: func() set.Set { return set.NewGoMap(0) }},
	"hashsha1":    {name: "HashSHA1", s: func() set.Set { return set.NewHashSHA1(0, true) }},
	"spooky128":   {name: "Spooky128", s: func() set.Set { return set.NewSpooky128(0, true) }},
	"farmhash128": {name: "Farmhash128", s: func() set.Set { return set.NewFarm128(0, true) }},
	"spooky64":    {name: "Spooky64", s: func() set.Set { return set.NewSpooky64(0, true) }},
	"farmhash64":  {name: "Farmhash64", s: func() set.Set { return set.NewFarm64(0, true) }},
	"ternary":     {name: "TernarySet", s: func() set.Set { return set.NewTernarySet() }},
	"tchappat":    {name: "TchapPatricia", s: func() set.Set { return set.NewTchapPatricia() }},
	"quicktrie":   {name: "Quicktrie", s: func() set.Set { return set.NewQuicktrie() }},
}

func decodeKeys(r io.Reader) (out []string, err error) {
	scan := bufio.NewScanner(r)
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		out = append(out, scan.Text())
	}
	err = scan.Err()
	return
}

func setImpl(settype string) (out []setimpl, err error) {

	if settype == "all" {
		for _, impl := range impls {
			out = append(out, impl)
		}
		return out, nil
	}

	impl, ok := impls[settype]
	if !ok {
		var keys []string
		for k := range impls {
			keys = append(keys, k)
		}

		return nil, fmt.Errorf("%q is not a valid set type, valids types are: %s",
			settype,
			strings.Join(keys, ", "))
	}

	out = append(out, impl)

	return
}
