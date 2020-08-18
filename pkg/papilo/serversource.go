package papilo

import (
	"context"
	"log"
	"net/http"
	"time"
)

// ServerSource implements a default server source
type ServerSource struct {
	srv *http.Server
	out chan []byte
}

// NewServerSource returns a server source.
// addr is the address and port to run the server on
func NewServerSource(addr string) ServerSource {
	srvSource := ServerSource{}
	srvSource.srv = &http.Server{
		Addr: addr,
	}
	srvSource.out = make(chan []byte, 1)
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
	go s.srv.ListenAndServe()
	for d := range s.out {
		if p.IsClosed {
			s.shutdown()
			break
		}
		p.Write(d)
	}
}

func (s ServerSource) shutdown() {
	log.Printf("server.shutdown: shutting down server on %s\n", s.srv.Addr)
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("server.shutdown: ", err)
	}
}
