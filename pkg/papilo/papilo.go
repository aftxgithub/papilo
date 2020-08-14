package papilo

import (
	"fmt"
	"sync"
)

// Papilo is the orchestrator for a pipeline
type Papilo struct {
	pipeline Pipeline
	wg       sync.WaitGroup
	running  bool
}

// New returns a new Papilo object
func New() *Papilo {
	pilo := Papilo{
		pipeline: newPipeline(),
		wg:       sync.WaitGroup{},
	}
	pilo.pipeline.sourcer = NewStdinSource()
	pilo.pipeline.sinker = NewStdoutSink()
	return &pilo
}

// Run starts the pipeline
func (p *Papilo) Run() error {
	if p.running {
		return fmt.Errorf("Pipeline already running")
	}

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

	// start the components
	// start readers first to prevent deadlock
	for i := hIndex; i >= 0; i-- {
		mchan := make(chan interface{})
		go p.pipeline.components[i](mchan, cchan)
		// Next component uses this channel as its output
		cchan = mchan
	}

	// start the source
	go p.pipeline.sourcer.Source(cchan)
	// pipeline is running
	p.running = true
	p.wg.Add(1)

	// block till pipeline is stopped
	p.wg.Wait()

	return nil
}

// Stop ends the pipeline
func (p *Papilo) Stop() {
	if p.running {
		p.wg.Done()
		p.running = false
	}
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
