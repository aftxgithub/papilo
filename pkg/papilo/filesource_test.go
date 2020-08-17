package papilo

import (
	"io/ioutil"
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
		Sourcer: NewFileSource(testFileName, ReadTypeLine),
		Sinker:  TestFileSourceSink{},
	}
	go p.Run(mains)

	time.Sleep(time.Second * 2)
	p.Stop()

	if fileSourceOutput != "Hello World!" {
		t.Errorf("Wrong data, expected 'Hello World!' got %s", fileSourceOutput)
	}
}

var count int

func countCmpnt(p *Pipe) {
	var cnt int
	for !p.IsClosed {
		_, err := p.Next()
		if err != nil {
			continue
		}
		cnt++
		p.Write(cnt)
	}
}

type testFileSourceCountSink struct{}

func (t testFileSourceCountSink) Sink(p *Pipe) {
	for !p.IsClosed {
		d, _ := p.Next()
		data, ok := d.(int)
		if !ok {
			continue
		}
		count = data
	}
}

func TestReadTypes(t *testing.T) {
	testFile, err := ioutil.TempFile("/tmp", "test*.txt")
	if err != nil {
		t.Error(err)
	}
	testFileName := testFile.Name()
	testFile.Write([]byte("Hello World!"))
	testFile.Close()

	p := New()
	mains := &Pipeline{
		Sourcer:    NewFileSource(testFileName, ReadTypeLine),
		Components: []Component{countCmpnt},
		Sinker:     testFileSourceCountSink{},
	}
	go p.Run(mains)
	time.Sleep(1 * time.Second)
	if count != 1 {
		t.Errorf("Expected count 1, got %d", count)
	}
	p.Stop()

	mains.Sourcer = NewFileSource(testFileName, ReadTypeWord)
	go p.Run(mains)
	time.Sleep(1 * time.Second)
	if count != 2 {
		t.Errorf("Expected count 2, got %d", count)
	}
	p.Stop()

	mains.Sourcer = NewFileSource(testFileName, ReadTypeByte)
	go p.Run(mains)
	time.Sleep(1 * time.Second)
	if count != 12 {
		t.Errorf("Expected count 12, got %d", count)
	}
	p.Stop()
}
