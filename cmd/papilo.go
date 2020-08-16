package main

import (
	"log"

	flag "github.com/spf13/pflag"
	"github.com/thealamu/papilo/internal/papilocmd"
	"github.com/thealamu/papilo/pkg/papilo"
)

var cfgFilePath string

func main() {
	parseFlags()

	cfg := papilocmd.Config(cfgFilePath)
	if cfg == nil {
		log.Println("Could not read config file, using defaults")
	}

	mains, err := papilocmd.BuildPipeline(cfg)
	if err != nil {
		log.Println(err)
		return
	}

	p := papilo.New()
	p.Run(mains)
}

func parseFlags() {
	flag.StringVarP(&cfgFilePath, "config", "c", "", "path to pipeline config file")
	flag.Parse()
}
