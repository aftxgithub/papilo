package papilo

import (
	"testing"
	"time"
)

type testSumSource struct{}

func (t testSumSource) Source(p *Pipe) {
	p.Write(float64(8))
	p.Write(float64(7))
}

var sumOutput float64

type testSumSink struct{}

func (t testSumSink) Sink(p *Pipe) {
	for {
		d, _ := p.Next()
		num, ok := d.(float64)
		if !ok {
			continue
		}
		sumOutput = num
	}

}

func TestSumComponent(t *testing.T) {
	p := New()
	mains := &Pipeline{
		Sourcer:    testSumSource{},
		Sinker:     testSumSink{},
		Components: []Component{SumComponent},
	}
	go p.Run(mains)

	time.Sleep(2 * time.Second)

	if sumOutput != 15 {
		t.Errorf("SumComponent does not do its work, expected 9, got %f", sumOutput)
	}
}
