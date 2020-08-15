package papilo

// Component is a worker in a pipeline.
// A component reads data from the in channel, does its work
// and writes data to the out channel
type Component func(*Pipe)
