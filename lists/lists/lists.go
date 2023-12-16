package lists

import "fmt"

type Product struct {
	id    int
	title string
	price float64
}

func lists() {
	// 1
	hobbies := [3]string{"Skateboarding", "Guitar", "Programming"}
	fmt.Println(hobbies)

	// 2
	first := hobbies[0]
	remaining := hobbies[1:]
	fmt.Println(first, remaining)

	// 3
	firstTwo := hobbies[:2]
	fmt.Println(firstTwo)

	firstTwoAlt := hobbies[0:2]
	fmt.Println(firstTwoAlt)

	// 4
	lastTwo := hobbies[1:]
	fmt.Println(lastTwo)

	// 5
	courseGoals := []string{"Learn to think in terms of Go", "Build a REST API"}
	fmt.Println(courseGoals)

	// 6
	courseGoals[1] = "Learn about Goroutines"
	courseGoals = append(courseGoals, "Build an app")
	fmt.Println(courseGoals)

	// 7
	products := []Product{
		{1, "first", 12.99},
		{2, "second", 20.49},
	}

	newProduct := Product{
		3,
		"third",
		14.90,
	}

	products = append([]Product(products), newProduct)

	fmt.Println(products)
}
