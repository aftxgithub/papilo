package papilo

import "fmt"

// Pipe implements a buffered pipe.
// A buffered pipe connects two filter components.
//
// A pipe buffers incoming data,
// pushes the next data to the requiring component
// and receives from another pipe.
// A pipe can also propagate a close request which closes all pipes in front of it.
type Pipe struct {
	// bufSize is the max data the pipe can hold at a time
	bufSize int
	// count is the number of data in the pipe at any time
	count int
	// buffer holds the pipe's data
	buffer []interface{}
	// out is the next pipe linked to this pipe
	out *Pipe
	// in is the channel through which a pipe receives data
	in chan interface{}
	// IsClosed is true only when a pipe has been closed
	// and is no more interested in data.
	IsClosed bool
}

func newPipe(bufSize int, next *Pipe) *Pipe {
	p := &Pipe{
		bufSize: bufSize,
		buffer:  make([]interface{}, bufSize),
		in:      make(chan interface{}),
		out:     next,
	}
	go p.listen()
	return p
}

// Next returns the next data in the buffer.
// An error is returned if there is no data in the buffer
func (p *Pipe) Next() (interface{}, error) {
	if p.count == 0 {
		return nil, fmt.Errorf("No data in buffer")
	}
	p.count--
	return p.buffer[p.count], nil
}

// Close closes a pipe and propagates the close to pipes after it to prevent a clog
func (p *Pipe) Close() {
	if p.out == nil {
		p.IsClosed = true
		return
	}
	p.out.Close() // propagate downstream
	p.IsClosed = true
}

// isFull returns true if buffer is full, false otherwise
func (p *Pipe) isFull() bool {
	return p.count == p.bufSize
}

// listen opens the pipe up for incoming data
func (p *Pipe) listen() {
	for {
		if p.IsClosed {
			return
		}
		if p.isFull() {
			continue
		}
		select {
		case data := <-p.in:
			// Insert in reverse order, end first
			ins := p.bufSize - p.count - 1
			p.buffer[ins] = data
			p.count++
		}
	}
}

// Write sends data to the next pipe
func (p *Pipe) Write(data interface{}) error {
	if p.out == nil {
		return fmt.Errorf("End of the line")
	}
	p.out.in <- data
	return nil
}
