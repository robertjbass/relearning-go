package recursion

import "fmt"

func factorial(number int) int {
	if number <= 1 {
		return number
	}

	return number * factorial(number-1)

}

func main() {
	result := factorial(5)
	fmt.Println(result)
}
