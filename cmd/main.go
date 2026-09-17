package main

import (
	"example.com/calculator/prices"
)

func main() {
	taxRate := []float64{0.0, 0.07, 0.1, 0.15}

	for _, tVal := range taxRate {
		priceJob := prices.NewTaxIncludedPriceJob(tVal)
		priceJob.Process()
	}

}
