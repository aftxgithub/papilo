package papilocmd

import (
	"os"
)

var cfgName = "pipeline.yaml"

// Cfg holds the pipeline configurations
type Cfg struct{}

// Config returns a new Cfg
func Config() (Cfg, error) {
	_, err := os.Stat(cfgName)
	if err != nil {
		return Cfg{}, err
	}
	return Cfg{}, nil
}
