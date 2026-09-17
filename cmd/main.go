package main

import (
	"fmt"

	"example.com/calculator/filemanager"
	"example.com/calculator/prices"
)

func main() {
	taxRate := []float64{0.0, 0.07, 0.1, 0.15}

	for _, tVal := range taxRate {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", tVal*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, tVal)
		priceJob.Process()
	}

}
