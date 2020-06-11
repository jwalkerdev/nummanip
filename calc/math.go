package calc

import "fmt"

// init - runs when the package is initialized. Can have multiple
func init() {
	fmt.Println("calc/math.go ==> init() [1]")
}

func init() {
	fmt.Println("calc/math.go ==> init() [2]")
}

// Add - returns sum of two integers
func Add(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum = sum + num
	}
	return sum
}
