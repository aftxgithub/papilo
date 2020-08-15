package main

import (
	"fmt"
	"log"

	flag "github.com/spf13/pflag"
	"github.com/thealamu/papilo/internal/papilocmd"
)

var cfgFilePath string

func main() {
	parseFlags()

	cfg := papilocmd.Config(cfgFilePath)
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
	flag.StringVarP(&cfgFilePath, "config", "c", "", "path to pipeline config file")
	flag.Parse()
}
