package papilo

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

type TestFileSinkSource struct{}

func (t TestFileSinkSource) Source(p *Pipe) {
	p.Write([]byte("Hello World!"))
}

func TestFileSink(t *testing.T) {
	testFile, err := ioutil.TempFile("/tmp", "test*.txt")
	if err != nil {
		t.Error(err)
	}
	testFileName := testFile.Name()
	testFile.Close()

	p := New()
	mains := &Pipeline{
		Sourcer: TestFileSinkSource{},
		Sinker:  NewFileSink(testFileName),
	}
	go p.Run(mains)

	time.Sleep(time.Second * 2)
	p.Stop()

	testFile, err = os.Open(testFileName)
	if err != nil {
		t.Error(err)
	}
	defer testFile.Close()

	data, err := ioutil.ReadFile(testFileName)
	if err != nil {
		t.Error(err)
	}

	if string(data) != "Hello World!" {
		t.Errorf("Wrong data, expected 'Hello World!', got %s", string(data))
	}
}
