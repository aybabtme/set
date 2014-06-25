package main

import (
	"github.com/aybabtme/set"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"runtime"
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "setbench"
	app.Usage = "Benchmarks different properties of set implementations."

	memplotFlags, memplotAction := memplotCommand()

	app.Commands = []cli.Command{
		{
			Name:   "memplot",
			Usage:  "Plots memory usage over time as keys are inserted in a set.",
			Flags:  memplotFlags,
			Action: memplotAction,
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
	s    set.Set
}

var impls = map[string]setimpl{
	"gomap":         {name: "GoMap", s: set.NewGoMap(0)},
	"hashsha1":      {name: "HashSHA1", s: set.NewHashSHA1(0)},
	"spooky128":     {name: "Spooky128", s: set.NewSpooky128(0)},
	"farmhash128":   {name: "Farmhash128", s: set.NewFarm128(0)},
	"spooky64":      {name: "Spooky64", s: set.NewSpooky64(0)},
	"farmhash64":    {name: "Farmhash64", s: set.NewFarm64(0)},
	"spooky32":      {name: "Spooky32", s: set.NewSpooky32(0)},
	"farmhash32":    {name: "Farmhash32", s: set.NewFarm32(0)},
	"adler32":       {name: "Adler32", s: set.NewAdler32(0)},
	"murmur32":      {name: "Murmur32", s: set.NewMurmur32(0)},
	"djb32":         {name: "Djb32", s: set.NewDjb32(0)},
	"elf32":         {name: "Elf32", s: set.NewElf32(0)},
	"java32":        {name: "Java32", s: set.NewJava32(0)},
	"jenkins32":     {name: "Jenkins32", s: set.NewJenkins32(0)},
	"sdbm32":        {name: "SDBM32", s: set.NewSDBM32(0)},
	"sqlite32":      {name: "SQLite32", s: set.NewSQLite32(0)},
	"superfasthash": {name: "SuperFastHash", s: set.NewSuperFastHash(0)},
}
