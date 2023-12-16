package main

import "fmt"

type floatMap map[string]float64

func main() {
	userNames := make([]string, 2, 5)
	userNames[0] = "Bob"
	userNames[1] = "Sam"

	fmt.Println(userNames, cap(userNames))

	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	fmt.Println(courseRatings)

	for index, value := range userNames {
		fmt.Println(index, value)
	}

	for key, value := range courseRatings {
		fmt.Println(key, value)
	}

}
