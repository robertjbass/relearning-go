package main

import (
	"errors"
	"fmt"
	"os"
)

func saveNumbersToFile(ebtString string, profitString string, ratioString string) {
	reportString := fmt.Sprintf("%v\n%v\n%v", ebtString, profitString, ratioString)
	fmt.Println(reportString)
	os.WriteFile("report.txt", []byte(reportString), 0644)
}

func main() {
	revenue, err := getUserInput("Revenue")
	if err != nil {
		panic(err)
	}
	expenses, err := getUserInput("Expenses")
	if err != nil {
		panic(err)
	}
	taxRate, err := getUserInput("Tax Rate")
	if err != nil {
		panic(err)
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	ebtString := fmt.Sprintf("EBT: $%.2f", ebt)
	profitString := fmt.Sprintf("Profit: $%.2f", profit)
	ratioString := fmt.Sprintf("Ratio: %.1f%%", ratio)

	saveNumbersToFile(ebtString, profitString, ratioString)
}

func getUserInput(prompt string) (float64, error) {
	var userInput float64
	fmt.Printf("%s: ", prompt)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		errorString := fmt.Sprintf("%s value must be higher than 0", prompt)
		return 0, errors.New(errorString)
	}
	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}
