package main

import "fmt"

func main() {
	age := 33

	fmt.Println("agePointer", &age)
	fmt.Println("age", age)
	getAdultYears(&age)
	fmt.Println("adultYears", age)

}

func getAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
