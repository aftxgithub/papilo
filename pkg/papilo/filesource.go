package papilo

import (
	"bufio"
	"io"
	"os"
)

// FileSource implements a default file data source
type FileSource struct {
	filepath string
	fdesc    *os.File
	bSize    int
}

// NewFileSource returns a new file data source for streaming bytes of a file.
// The path parameter is the path of the file to be read,
// byteSize is the number of bytes to write out at a time.
func NewFileSource(path string, byteSize int) FileSource {
	return FileSource{
		filepath: path,
		bSize:    byteSize,
	}
}

// NewFdSource returns a new file data source for streaming bytes of a file.
// The fd parameter is an opened file to be read,
// byteSize is the number of bytes to write out at a time.
func NewFdSource(fd *os.File, byteSize int) FileSource {
	return FileSource{
		fdesc: fd,
		bSize: byteSize,
	}
}

// Source is the file implementation for the Sourcer interface.
// Defined output for this source is a slice of bytes.
func (f FileSource) Source(p *Pipe) {
	var fd *os.File = f.fdesc
	if fd == nil {
		var err error
		fd, err = os.Open(f.filepath)
		if err != nil {
			panic(err)
		}
	}

	reader := bufio.NewReader(fd)
	for {
		dataBuf := make([]byte, f.bSize)
		_, err := reader.Read(dataBuf)
		if err == io.EOF {
			break
		}
		p.Write(dataBuf)
	}
}
