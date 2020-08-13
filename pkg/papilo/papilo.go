package papilo

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
