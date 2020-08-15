package main

import (
	"fmt"
	"log"

	flag "github.com/spf13/pflag"
	"github.com/thealamu/papilo/internal/papilocmd"
)

var cfgFileName string

func main() {
	parseFlags()

	cfg := papilocmd.Config(cfgFileName)
	if cfg == nil {
		fmt.Println("Could not read config file, using default")
	}

	p, err := papilocmd.Build(cfg)
	if err != nil {
		log.Println(err)
		return
	}

	p.Run()
}

func parseFlags() {
	flag.StringVarP(&cfgFileName, "config", "c", "", "path to pipeline config file")
	flag.Parse()
}
