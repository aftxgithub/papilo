package papilo

import "fmt"

// Papilo is the orchestrator for a pipeline
type Papilo struct {
	pipeline Pipeline
}

// New returns a new Papilo object
func New() Papilo {
	return Papilo{
		pipeline: newPipeline(),
	}
}

// Run starts the pipeline
func (p Papilo) Run() error {
	if p.pipeline.sourcer == nil {
		return fmt.Errorf("Data source not defined")
	}
	if p.pipeline.sinker == nil {
		return fmt.Errorf("Data sink not defined")
	}

	hIndex := len(p.pipeline.components) - 1
	cchan := make(chan []byte)

	// start the sink
	go p.pipeline.sinker.Sink(cchan)

	// start the components, readers first
	for i := hIndex; i >= 0; i-- {
		mchan := make(chan []byte)
		go p.pipeline.components[i](mchan, cchan)
		// Next component uses this channel as its output
		cchan = mchan
	}

	// start the source
	go p.pipeline.sourcer.Source(cchan)

	return nil
}

// SetSource registers a data source for the pipeline
func (p Papilo) SetSource(s Sourcer) {
	p.pipeline.sourcer = s
}

// SetSink registers a data sink for the pipeline
func (p Papilo) SetSink(s Sinker) {
	p.pipeline.sinker = s
}

// AddComponent adds a component to the pipeline
func (p Papilo) AddComponent(c Component) {
	p.pipeline.components = append(p.pipeline.components, c)
}
