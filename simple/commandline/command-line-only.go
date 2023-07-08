package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/uniyuni1101/studycfg/simple/config"
)

var CurrentConfig = &config.Config{}

type CommandLineConfigLoader struct {
	flagSet       flag.FlagSet
	defaultConfig config.Config
}

func main() {
	f := flag.NewFlagSet("single", flag.ExitOnError)

	f.StringVar((*string)(&CurrentConfig.Theme), "theme", "simple", "select display theme")
	f.StringVar((*string)(&CurrentConfig.Theme), "t", "simple", "select display theme (shorthand)")

	f.IntVar((*int)(&CurrentConfig.TickRate), "rate", 20, "change display update per second")
	f.IntVar((*int)(&CurrentConfig.TickRate), "r", 20, "change display update per second (shorthand)")

	f.Parse(os.Args[1:])

	fmt.Printf("%#+v\n", CurrentConfig)
}
