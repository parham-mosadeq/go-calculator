package conversion

import (
	"errors"
	"strconv"
)

func StringToFloats(strings []string) ([]float64, error) {
	var floats []float64

	for _, lVal := range strings {
		floatPrice, err := strconv.ParseFloat(lVal, 64)
		if err != nil {
			return nil, errors.New("Something went wrong ")
		}
		floats = append(floats, floatPrice)
	}
	return floats, nil
}
