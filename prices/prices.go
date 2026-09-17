package prices

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)

	for _, pVal := range job.InputPrices {
		result[fmt.Sprint("%.2f", pVal)] = pVal * (1 + job.TaxRate)
	}

	fmt.Println("Processed tax rates & prices", result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		InputPrices: []float64{10.0, 20.0, 30.0},
		TaxRate:     taxRate,
	}

}
