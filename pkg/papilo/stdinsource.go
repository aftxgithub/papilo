package papilo

import "os"

// NewStdinSource returns a new file data source that reads from standard input
func NewStdinSource() FileSource {
	return NewFdSource(os.Stdin)
}
