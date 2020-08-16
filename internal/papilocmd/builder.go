package papilocmd

import (
	"fmt"

	"github.com/thealamu/papilo/pkg/papilo"
)

// BuildPipeline builds a papilo pipeline using the config
func BuildPipeline(cfg *Cfg) (*papilo.Pipeline, error) {
	mains := &papilo.Pipeline{}

	// Set the source
	err := buildSource(cfg.Pipeline.Source, mains)
	if err != nil {
		return nil, err
	}

	// Add components
	err = buildCmpnts(cfg.Pipeline.Components, mains)
	if err != nil {
		return nil, err
	}

	// Set the sink
	err = buildSink(cfg.Pipeline.Sink, mains)
	if err != nil {
		return nil, err
	}

	return mains, nil
}

func buildCmpnts(cmpnts []string, pline *papilo.Pipeline) (err error) {
	for _, cmpnt := range cmpnts {
		switch cmpnt {
		case "sum":
			pline.Components = append(pline.Components, papilo.SumComponent)
		default:
			err = fmt.Errorf("Invalid component %s", cmpnt)
		}
	}
	return
}

func buildSource(src source, pline *papilo.Pipeline) (err error) {
	switch src.Type {
	case "file":
		err = setFileSource(pline, src.Config)
	case "stdin":
		// do nothing, source defaults to stdin
	default:
		err = fmt.Errorf("Invalid source %s", src.Type)
	}
	return
}

func setFileSource(p *papilo.Pipeline, config map[string]interface{}) error {
	pI := config["path"]
	path, ok := pI.(string)
	if !ok {
		return fmt.Errorf("Source filepath should be string")
	}
	p.Sourcer = papilo.NewFileSource(path)
	return nil
}

func buildSink(snk sink, pline *papilo.Pipeline) (err error) {
	switch snk.Type {
	case "file":
		err = setFileSink(pline, snk.Config)
	case "stdout":
		// do nothing, defaults to stdout
	default:
		err = fmt.Errorf("Invalid sink %s", snk.Type)
	}
	return
}

func setFileSink(p *papilo.Pipeline, config map[string]interface{}) error {
	pI := config["path"]
	path, ok := pI.(string)
	if !ok {
		return fmt.Errorf("Sink filepath should be string")
	}
	p.Sinker = papilo.NewFileSink(path)
	return nil
}
