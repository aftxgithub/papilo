package papilo

// Sinker defines methods for any data sink.
// A data sink should implement storage for processed data
type Sinker interface {
	// Sink implements storage for processed data
	// Sink reads from the provided channel
	Sink(*Pipe)
}
