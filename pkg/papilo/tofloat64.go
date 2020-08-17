package papilo

import (
	"fmt"
	"strconv"
)

// ToFloat64 transforms data to float64
func ToFloat64(data interface{}) (float64, error) {
	d, ok := data.(float64)
	if ok {
		return d, nil
	}

	switch v := data.(type) {
	case int:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case string:
		return strconv.ParseFloat(v, 64)
	case []byte:
		return strconv.ParseFloat(string(v), 64)
	default:
		return 0, fmt.Errorf("Could not convert to float64")
	}
}
