package papilo

import "net/http"

// ServerSource implements a default server source
type ServerSource struct {
	s   *http.Server
	out chan []byte
}

// NewServerSource returns a server source.
// addr is the address and port to run the server on
func NewServerSource(addr string) ServerSource {
	srvSource := ServerSource{}
	srvSource.s = &http.Server{
		Addr: addr,
	}
	srvSource.out = make(chan []byte)
	setHandler(&srvSource)
	return srvSource
}

func setHandler(s *ServerSource) {
	http.Handle("/", s)
}

func (s ServerSource) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

// Source is the server implementation for the Sourcer interface.
// Defined output for this source is a slice of bytes.
func (s ServerSource) Source(p *Pipe) {

}
