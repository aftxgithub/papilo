package papilo

// FileSource implements a default file data source
type FileSource struct{}

// NewFileSource returns a new file data source for streaming lines of a file.
// The path parameter is the path of the file to be read.
func NewFileSource(path string) FileSource {
	return FileSource{}
}

func (f FileSource) source(out chan interface{}) {

}
