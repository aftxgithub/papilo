package main

import (
	"fmt"

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
}

func parseFlags() {
	flag.StringVarP(&cfgFileName, "config", "c", "", "path to pipeline config file")
	flag.Parse()
}
