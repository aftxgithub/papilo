package papilo

import "os"

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
