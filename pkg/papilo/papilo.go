package papilo

import (
	"fmt"
	"sync"
)

// Papilo is the orchestrator for a pipeline
type Papilo struct {
	wg      sync.WaitGroup
	running bool
}

// New returns a new Papilo object
func New() *Papilo {
	pilo := Papilo{
		wg: sync.WaitGroup{},
	}
	return &pilo
}

// Run starts the pipeline
func (p *Papilo) Run(pline *Pipeline) error {
	if p.running {
		return fmt.Errorf("Pipeline already running")
	}

	if pline == nil {
		return fmt.Errorf("Can't start nil pipeline")
	}
	if pline.Sourcer == nil {
		pline.Sourcer = NewStdinSource()
	}
	if pline.Sinker == nil {
		pline.Sinker = NewStdoutSink()
	}
	if pline.BufSize <= 0 {
		pline.BufSize = 1
	}

	hIndex := len(pline.Components) - 1
	//cchan := make(chan interface{})
	cpipe := newPipe(pline.BufSize, nil)

	// start the sink
	go pline.Sinker.Sink(cpipe)

	// start the components
	// start readers first to prevent deadlock
	for i := hIndex; i >= 0; i-- {
		mpipe := newPipe(pline.BufSize, cpipe)
		go pline.Components[i](mpipe)
		// Next component uses this channel as its output
		cpipe = mpipe
	}

	// start the source
	spipe := newPipe(0, cpipe)
	go pline.Sourcer.Source(spipe)
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
