package papilo

import "testing"

func TestToBytes(t *testing.T) {
	val, err := ToBytes(5.9)
	if err != nil {
		t.Error(err)
	}

	if string(val) != "5.9" {
		t.Errorf("Wrong bytes result, expected %s, got %s", "5.9", string(val))
	}
}
