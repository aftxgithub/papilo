package papilocmd

import (
	"fmt"

	"github.com/thealamu/papilo/pkg/papilo"
)

// Build builds a papilo pipeline using the config
func Build(cfg *Cfg) (*papilo.Papilo, error) {
	p := papilo.New()
	if cfg == nil {
		return p, nil
	}

	var err error

	// Set the source
	switch cfg.Pipeline.Source.Type {
	case "file":
		err = setFileSource(p, cfg.Pipeline.Source.Config)
	case "stdin":
		// do nothing, source defaults to stdin
	default:
		err = fmt.Errorf("Invalid source %s", cfg.Pipeline.Source.Type)
	}

	// Add components
	for _, cpnt := range cfg.Pipeline.Components {
		switch cpnt {
		case "sum":
			addSumComponent(p)
		}
	}

	// Set the sink
	switch cfg.Pipeline.Sink.Type {
	case "file":
		err = setFileSink(p, cfg.Pipeline.Sink.Config)
	case "stdout":
		// do nothing, defaults to stdout
	default:
		err = fmt.Errorf("Invalid sink %s", cfg.Pipeline.Sink.Type)
	}

	return p, err
}

func setFileSource(p *papilo.Papilo, config map[string]interface{}) error {
	pI := config["path"]
	path, ok := pI.(string)
	if !ok {
		return fmt.Errorf("Source filepath should be string")
	}
	p.SetSource(papilo.NewFileSource(path))
	return nil
}

func addSumComponent(p *papilo.Papilo) {
	p.AddComponent(papilo.SumComponent)
}

func setFileSink(p *papilo.Papilo, config map[string]interface{}) error {
	pI := config["path"]
	path, ok := pI.(string)
	if !ok {
		return fmt.Errorf("Sink filepath should be string")
	}
	p.SetSink(papilo.NewFileSink(path))
	return nil
}
