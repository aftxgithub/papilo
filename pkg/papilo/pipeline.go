package papilo

// Pipeline represents a whole pipeline of components
type Pipeline struct {
	sourcer    Sourcer
	components []Component
	sinker     Sinker
}
