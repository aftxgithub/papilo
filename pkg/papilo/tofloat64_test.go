package papilo

import "testing"

func TestToFloat64(t *testing.T) {
	val, err := ToFloat64([]byte("5"))
	if err != nil {
		t.Error(err)
	}
	if val != 5 {
		t.Errorf("Unexpected float result, expected %d, got %f", 5, val)
	}
}
