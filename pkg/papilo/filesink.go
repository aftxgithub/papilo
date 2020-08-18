package papilo

import (
	"log"
	"os"
)

// FileSink implements a default file data sink
type FileSink struct {
	filepath string
	fdesc    *os.File
}

// NewFileSink returns a new file data sink for writing lines to a file.
// The path parameter is the path of the file to be written.
func NewFileSink(path string) FileSink {
	return FileSink{
		filepath: path,
	}
}

// NewFdSink returns a new file data sink for writing lines to a file.
// The fd parameter is an opened file to be written.
func NewFdSink(fd *os.File) FileSink {
	return FileSink{
		fdesc: fd,
	}
}

// Sink is the implementation for the Sinker interface.
// Defined input for this sink is a slice of bytes.
// Sink will create a file if it does not exist
// or truncate the file otherwise.
func (f FileSink) Sink(p *Pipe) {
	var fd *os.File = f.fdesc
	if fd == nil {
		var err error
		fd, err = os.Create(f.filepath)
		if err != nil {
			panic(err)
		}
	}
	defer fd.Close()

	for !p.IsClosed {
		d, err := p.Next()
		if err != nil {
			// no data in pipe
			continue
		}
		data, err := ToBytes(d)
		if err != nil {
			log.Println(err)
			continue
		}
		fd.Write(data)
		fd.Write([]byte("\n"))
	}
}
