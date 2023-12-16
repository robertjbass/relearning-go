package main

import "fmt"

func main() {
	sum := add(1, 5)

	fmt.Println(sum)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
