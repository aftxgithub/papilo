package papilo

import (
	"bufio"
	"os"
)

// FileSource implements a default file data source
type FileSource struct {
	filepath string
}

// NewFileSource returns a new file data source for streaming lines of a file.
// The path parameter is the path of the file to be read.
func NewFileSource(path string) FileSource {
	return FileSource{
		filepath: path,
	}
}

// Source is the implementation for the Sourcer interface.
// Defined output for this source is a slice of bytes.
func (f FileSource) Source(out chan interface{}) {
	fd, err := os.Open(f.filepath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		out <- scanner.Bytes()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
