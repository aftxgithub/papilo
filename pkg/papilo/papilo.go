package papilo

// Papilo is the orchestrator for a pipeline
type Papilo struct {
	pipeline Pipeline
}

// New returns a new Papilo object
func New() Papilo {
	return Papilo{}
}
