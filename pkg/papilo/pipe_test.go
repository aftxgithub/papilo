package papilo

import (
	"testing"
	"time"
)

func TestPipeNext(t *testing.T) {
	testPipe := newPipe(1, nil)
	prevPipe := newPipe(1, testPipe)

	_, err := testPipe.Next()
	if err == nil {
		t.Errorf("Expected empty buffer error, got nil")
	}

	prevPipe.Write(5)
	time.Sleep(1 * time.Second)
	data, err := testPipe.Next()
	if err != nil {
		t.Error(err)
	}
	d, ok := data.(int)
	if !ok {
		t.Errorf("Unexpected type, expected int")
		return
	}
	if d != 5 {
		t.Errorf("Invalid data, expected %d, got %d", 5, d)
	}
}

func TestPipeFull(t *testing.T) {
	testPipe := newPipe(1, nil)
	prevPipe := newPipe(1, testPipe)

	prevPipe.Write(5)
	time.Sleep(2 * time.Second)
	if !testPipe.isFull() {
		t.Errorf("Pipe does not know when it is full")
	}
}

func TestPipeClose(t *testing.T) {
	nextPipe := newPipe(2, nil)
	testPipe := newPipe(2, nextPipe)
	prevPipe := newPipe(2, testPipe)

	// propagate a close
	prevPipe.Close()

	if !nextPipe.IsClosed {
		t.Errorf("Pipe close does not propagate")
	}
}

func TestPipe(t *testing.T) {
	nextPipe := newPipe(2, nil)
	testPipe := newPipe(2, nextPipe)
	prevPipe := newPipe(2, testPipe)

	prevPipe.Write(6)
	prevPipe.Write(5)
	time.Sleep(2 * time.Second) // wait for the receiver to read so we can do checks
	if testPipe.count != 2 {
		t.Errorf("Listen does not increment count, expected %d, got %d", 2, testPipe.count)
	}
	sd, ok := testPipe.buffer[1].(int)
	if !ok || sd != 6 {
		t.Errorf("Write method does not send data to next pipe")
	}

	err := nextPipe.Write(5)
	if err == nil {
		t.Errorf("Expected end of the line error, got nil")
	}
}
