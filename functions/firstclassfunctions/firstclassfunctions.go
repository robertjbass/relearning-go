package firstclassfunctions

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	updatedNumbers := []int{}
	for _, value := range *numbers {
		updatedNumbers = append(updatedNumbers, transform(value))
	}

	return updatedNumbers
}
