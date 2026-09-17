package prices

import (
	"fmt"

	"example.com/calculator/conversion"
	"example.com/calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadDate() {
	// file, err := os.Open("prices.txt")
	// if err != nil {
	// 	fmt.Println("Reading file failed!", err)
	// 	return
	// }

	// scanner := bufio.NewScanner(file)

	// var lines []string
	lines, err := filemanager.ReadLines("prices.txt")

	if err != nil {
		fmt.Println("Reading the file content failed!", err)
		return
	}

	// for scanner.Scan() {
	// 	lines = append(lines, scanner.Text())
	// }

	// err = scanner.Err()
	// if err != nil {
	// 	fmt.Println("Reading the file content failed!", err)
	// 	file.Close()
	// 	return
	// }

	// prices := make([]float64, len(lines))
	prices, err := conversion.StringToFloats(lines)

	if err != nil {
		fmt.Println("Reading prices failed", err)
	}

	// for lIdx, lVal := range lines {
	// 	floatPrice, err := strconv.ParseFloat(lVal, 64)
	// 	if err != nil {
	// 		file.Close()
	// 		fmt.Println("Reading prices failed", err)
	// 	}

	// 	prices[lIdx] = floatPrice
	// }

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadDate()
	result := make(map[string]string)

	for _, pVal := range job.InputPrices {
		taxIncludedPrice := pVal * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", pVal)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println("Processed tax rates and prices = ", result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		InputPrices: []float64{10.0, 20.0, 30.0},
		TaxRate:     taxRate,
	}

}
