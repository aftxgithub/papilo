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
