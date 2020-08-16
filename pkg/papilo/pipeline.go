package papilo

// Pipeline represents a whole pipeline of components
type Pipeline struct {
	// Data Source for pipeline
	Sourcer Sourcer

	// Components for pipeline, order matters
	Components []Component

	// Data Sink for pipeline
	Sinker Sinker

	// BufSize is the size of a pipe buffer
	BufSize int
}
