package papilo

import "fmt"

// ToBytes converts data to byte slice
func ToBytes(data interface{}) ([]byte, error) {
	d, ok := data.([]byte)
	if ok {
		return d, nil
	}

	switch v := data.(type) {
	case string:
		return []byte(v), nil
	default:
		return []byte(fmt.Sprint(data)), nil
	}
}
