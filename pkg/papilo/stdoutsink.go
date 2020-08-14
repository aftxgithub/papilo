package papilo

import "os"

// NewStdoutSink returns a new file data sink that writes to standard output
func NewStdoutSink() FileSink {
	return NewFdSink(os.Stdout)
}
