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
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	const useFileMananger = true
	const inputFilePath = "prices.txt"

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)

		var manager iomanager.IOManager

		if useFileMananger {
			manager = filemanager.New(inputFilePath, fmt.Sprintf("result_%.0f.json", taxRate*100))
		} else {
			manager = cmdmanager.New()
		}

		priceJob := *prices.NewTaxIncludedPriceJob(manager, taxRate)

		go priceJob.Process(doneChans[index], errorChans[index])
	}

	for index := range taxRates {
		// whichever channel emits a value first will be processed
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done")
		}
	}

}
