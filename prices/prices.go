package prices

import (
	"fmt"

	"example.com/calculator/conversion"
	"example.com/calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64           `json:"tax_rate"`
	InputPrices       []float64         `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_include_prices"`
	IoMnanager        filemanager.FileManager
}

func (job *TaxIncludedPriceJob) LoadDate() {
	// file, err := os.Open("prices.txt")
	// if err != nil {
	// 	fmt.Println("Reading file failed!", err)
	// 	return
	// }

	// scanner := bufio.NewScanner(file)

	// var lines []string
	lines, err := job.IoMnanager.ReadLines()

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

	job.TaxIncludedPrices = result
	err := job.IoMnanager.WriteJSON(job)
	if err != nil {

		fmt.Println("saving failed")
	}

}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		InputPrices: []float64{10.0, 20.0, 30.0},
		TaxRate:     taxRate,
		IoMnanager:  fm,
	}

}
