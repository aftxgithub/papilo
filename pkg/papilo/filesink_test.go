package papilo

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

type TestFileSinkSource struct{}

func (t TestFileSinkSource) Source(out chan interface{}) {
	out <- []byte("Hello World!")
}

func TestFileSink(t *testing.T) {
	testFile, err := ioutil.TempFile("/tmp", "test*.txt")
	if err != nil {
		t.Error(err)
	}
	testFileName := testFile.Name()
	testFile.Close()

	p := New()
	p.SetSource(TestFileSinkSource{})
	p.SetSink(NewFileSink(testFileName))

	go p.Run()
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
