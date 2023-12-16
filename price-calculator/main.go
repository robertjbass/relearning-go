package main

import (
	"fmt"

	"github.com/robertjbass/price-calculator/cmdmanager"
	"github.com/robertjbass/price-calculator/filemanager"
	"github.com/robertjbass/price-calculator/iomanager"
	"github.com/robertjbass/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	const useFileMananger = true
	const inputFilePath = "prices.txt"

	for _, taxRate := range taxRates {

		var manager iomanager.IOManager

		if useFileMananger {
			manager = filemanager.New(inputFilePath, fmt.Sprintf("result_%.0f.json", taxRate*100))
		} else {
			manager = cmdmanager.New()
		}

		priceJob := *prices.NewTaxIncludedPriceJob(manager, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("could not process job")
			fmt.Println(err)
		}
	}

}
