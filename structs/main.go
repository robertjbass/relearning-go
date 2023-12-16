package main

import (
	"fmt"

	"robertjbass/structs/user"
)

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

func main() {
	userFirstName := getUserData("First name: ")
	userLastName := getUserData("Last name: ")
	userBirthdate := getUserData("Birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("user@mailinator.com", "Testing123")
	admin.Describe()

	appUser.Describe()
	appUser.ClearUserName()
	appUser.Describe()
}
