package papilocmd

import (
	"os"
)

var cfgName = "pipeline.yaml"

// Cfg holds the pipeline configurations
type Cfg struct {
	Pipeline pipeline `yaml:"pipeline"`
}

type pipeline struct {
	Source     source   `yaml:"source"`
	Components []string `yaml:"components"`
	Sink       string   `yaml:"sink"`
}

type source struct {
	Type   string                 `yaml:"type"`
	Config map[string]interface{} `yaml:"config"`
}

type sink struct {
	Type   string                 `yaml:"type"`
	Config map[string]interface{} `yaml:"config"`
}

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
