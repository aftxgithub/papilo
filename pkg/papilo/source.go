package papilo

// Sourcer defines methods for a data source
type Sourcer interface {
	// Source implements getting data and streaming into the provided channel
	Source(out chan []byte)
}
