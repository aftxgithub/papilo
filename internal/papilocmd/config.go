package papilocmd

import (
	"os"

	"gopkg.in/yaml.v2"
)

var cfgName = "pipeline.yaml"

// Cfg holds the pipeline configurations
type Cfg struct {
	Pipeline pipeline `yaml:"pipeline"`
	path     string
}

type pipeline struct {
	Source     source   `yaml:"source"`
	Components []string `yaml:"components"`
	Sink       sink     `yaml:"sink"`
}

type source struct {
	Type   string                 `yaml:"type"`
	Config map[string]interface{} `yaml:"config"`
}

type sink struct {
	Type   string                 `yaml:"type"`
	Config map[string]interface{} `yaml:"config"`
}

// Config returns the read config
func Config(filepath string) *Cfg {
	if filepath == "" {
		filepath = cfgName
	}

	fd, err := os.Open(filepath)
	if err != nil {
		return nil
	}
	defer fd.Close()

	var cfg Cfg
	err = yaml.NewDecoder(fd).Decode(&cfg)
	if err != nil {
		return nil
	}

	return &cfg
}
