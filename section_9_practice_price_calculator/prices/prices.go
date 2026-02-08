package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type taxIncludedPriceJob struct {
	IOManager         filemanager.FileManager `json:"-"`
	TaxRate           float64                 `json:"tax_rate"`
	InputPrices       []float64               `json:"input_prices"`
	TaxIncludedPrices map[string]string       `json:"tax_included_prices"`
}

func (job *taxIncludedPriceJob) LoadData() {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringToFLoats(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *taxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrices = result

	job.IOManager.WriteResult(job)

}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *taxIncludedPriceJob {
	return &taxIncludedPriceJob{
		IOManager:   fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
