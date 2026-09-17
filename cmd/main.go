package main

import "fmt"

func main() {
	prices := []float64{10.0, 20.0, 30.0}
	taxRate := []float64{0.0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, tVal := range taxRate {
		taxIncludePrice := make([]float64, len(prices))
		for pIdx, pVal := range prices {
			taxIncludePrice[pIdx] = pVal * (1 + tVal)
		}

		result[tVal] = taxIncludePrice
	}

	fmt.Println(result)
}
