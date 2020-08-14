package papilo

import (
	"testing"
	"time"
)

type testSource struct{}

func (t testSource) Source(out chan interface{}) {
	out <- 5
}

var output int

type testSink struct{}

func (t testSink) Sink(in chan interface{}) {
	for d := range in {
		var ok bool
		output, ok = d.(int)
		if !ok {
			panic("Expected data type int in sink")
		}
	}
}

func TestIntegration(t *testing.T) {
	p := New()
	p.SetSource(testSource{})
	p.SetSink(testSink{})
	p.AddComponent(func(in chan interface{}, out chan interface{}) {
		for d := range in {
			data, ok := d.(int)
			if !ok {
				t.Errorf("Expected data type int in squaring component")
				return
			}
			o := data * data
			out <- o
		}
	})
	go p.Run()
	time.Sleep(2 * time.Second)
	p.Stop()

	if output != 25 {
		t.Errorf("Wrong value, expected 25, got %d", output)
	}
}
