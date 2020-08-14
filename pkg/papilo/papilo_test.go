package papilo

import (
	"testing"
	"time"
)

type testSource struct{}

func (t testSource) Source(out chan []byte) {
	out <- []byte{5}
}

var output int

type testSink struct{}

func (t testSink) Sink(in chan []byte) {
	for d := range in {
		output = int(d[0])
	}
}

func TestIntegration(t *testing.T) {
	p := New()
	p.SetSource(testSource{})
	p.SetSink(testSink{})
	p.AddComponent(func(in chan []byte, out chan []byte) {
		for d := range in {
			data := int(d[0])
			o := data * data
			out <- []byte(string(o))
		}
	})
	go p.Run()
	time.Sleep(2 * time.Second)
	p.Stop()

	if output != 25 {
		t.Errorf("Wrong value, expected 25, got %d", output)
	}
}
