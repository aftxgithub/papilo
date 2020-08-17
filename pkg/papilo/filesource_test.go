package papilo

import (
	"io/ioutil"
	"log"
	"testing"
	"time"
)

var fileSourceOutput string

type TestFileSourceSink struct{}

func (t TestFileSourceSink) Sink(p *Pipe) {
	for !p.IsClosed {
		d, err := p.Next()
		if err != nil {
			// no data
			continue
		}
		intmed, ok := d.([]byte)
		if !ok {
			panic("Expected string data in TestFileSourceSink")
		}
		log.Println(intmed)
		fileSourceOutput = string(intmed)
	}
}

func TestFileSource(t *testing.T) {
	testFile, err := ioutil.TempFile("/tmp", "test*.txt")
	if err != nil {
		t.Error(err)
	}
	testFileName := testFile.Name()
	testFile.Write([]byte("Hello World!"))
	testFile.Close()

	p := New()
	mains := &Pipeline{
		Sourcer: NewFileSource(testFileName, 32),
		Sinker:  TestFileSourceSink{},
	}
	go p.Run(mains)

	time.Sleep(time.Second * 2)
	p.Stop()

	if fileSourceOutput != "Hello World!" {
		t.Errorf("Wrong data, expected 'Hello World!' got %s", fileSourceOutput)
	}
}
