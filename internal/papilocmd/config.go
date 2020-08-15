package papilocmd

import (
	"os"
)

var cfgName = "pipeline.yaml"

// Cfg holds the pipeline configurations
type Cfg struct{}

// Config returns a new Cfg
func Config(filename string) *Cfg {
	if filename == "" {
		filename = cfgName
	}
	_, err := os.Stat(filename)
	if err != nil {
		return nil
	}
	return &Cfg{}
}
