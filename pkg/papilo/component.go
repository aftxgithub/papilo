package papilo

// Component is a worker in a pipeline.
// A component reads data from the pipe, does its work
// and writes data to the next pipe
type Component func(*Pipe)
