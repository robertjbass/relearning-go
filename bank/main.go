package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/robertjbass/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println(randomdata.PhoneNumber())

	var accountBalance, err = fileops.ReadFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("---ERROR---")
		fmt.Println(err)
		fmt.Println("-----------")
	}

	for {
		choice := promptUser()

		switch choice {
		case 1:
			fmt.Printf("Your balance is %.2f\n", accountBalance)

		case 2:
			fmt.Print("How much do you want to deposit? ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit must be more than 0")
				continue
			}

			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)

			fmt.Println("Balance updated! New amount:", accountBalance)
		case 3:
			fmt.Print("How much do you want to withdraw? ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds, current balance:", accountBalance)
				continue
			}

			accountBalance -= withdrawalAmount
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)

			fmt.Println("Balance updated! New amount:", accountBalance)
		case 4:
			fmt.Println("Thank you")
			return

		default:
			fmt.Println("Invalid Input, try again")

		}

	}

}

func promptUser() int {
	presentOptions()

	var choice int
	fmt.Scan(&choice)
	return choice
}
