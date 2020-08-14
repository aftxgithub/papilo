package papilo

import (
	"fmt"
	"sync"
)

// Papilo is the orchestrator for a pipeline
type Papilo struct {
	pipeline Pipeline
	wg       sync.WaitGroup
}

// New returns a new Papilo object
func New() Papilo {
	return Papilo{
		pipeline: newPipeline(),
		wg:       sync.WaitGroup{},
	}
}

// Run starts the pipeline
func (p *Papilo) Run() error {
	if p.pipeline.sourcer == nil {
		return fmt.Errorf("Data source not defined")
	}
	if p.pipeline.sinker == nil {
		return fmt.Errorf("Data sink not defined")
	}

	hIndex := len(p.pipeline.components) - 1
	cchan := make(chan interface{})

	// start the sink
	go p.pipeline.sinker.Sink(cchan)

	// start the components, readers first
	for i := hIndex; i >= 0; i-- {
		mchan := make(chan interface{})
		go p.pipeline.components[i](mchan, cchan)
		// Next component uses this channel as its output
		cchan = mchan
	}

	// start the source
	go p.pipeline.sourcer.Source(cchan)
	p.wg.Add(1)

	// block till pipeline is stopped
	p.wg.Wait()

	return nil
}

// Stop ends the pipeline
func (p *Papilo) Stop() {
	p.wg.Done()
}

// SetSource registers a data source for the pipeline
func (p *Papilo) SetSource(s Sourcer) {
	p.pipeline.sourcer = s
}

// SetSink registers a data sink for the pipeline
func (p *Papilo) SetSink(s Sinker) {
	p.pipeline.sinker = s
}

// AddComponent adds a component to the pipeline
func (p *Papilo) AddComponent(c Component) {
	p.pipeline.components = append(p.pipeline.components, c)
}
