package papilo

import (
	"testing"
	"time"
)

type testSource struct{}

func (t testSource) Source(p *Pipe) {
	p.Write(5)
}

var output int

type testSink struct{}

func (t testSink) Sink(p *Pipe) {
	for !p.IsClosed {
		var ok bool
		d, err := p.Next()
		if err != nil {
			continue
		}
		output, ok = d.(int)
		if !ok {
			panic("Expected data type int in sink")
		}
	}
}

func TestIntegration(t *testing.T) {
	p := New()

	squareCmpt := func(p *Pipe) {
		for !p.IsClosed {
			d, err := p.Next()
			if err != nil {
				continue
			}
			data, ok := d.(int)
			if !ok {
				t.Errorf("Expected data type int in squaring component")
				return
			}
			o := data * data
			p.Write(o)
		}
	}

	mains := &Pipeline{
		Sourcer:    testSource{},
		Sinker:     testSink{},
		Components: []Component{squareCmpt},
	}
	go p.Run(mains)

	time.Sleep(2 * time.Second)
	p.Stop()

	if output != 25 {
		t.Errorf("Wrong value, expected 25, got %d", output)
	}
}
