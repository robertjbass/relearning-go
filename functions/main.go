package main

import "fmt"

func main() {

	sum := sumLessFirst(1, 2, 3)
	fmt.Println(sum)

	numbers := []int{1, 2, 3}
	anotherSum := sumLessFirst(0, numbers...)
	fmt.Println(anotherSum)

}

// variadic functions
func sumLessFirst(firstValue int, remaining ...int) int {
	fmt.Println("removing", firstValue)
	sum := 0

	for _, val := range remaining {
		sum += val
	}

	return sum
}
