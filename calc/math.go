package calc

import "fmt"
import "errors"

// init - runs when the package is initialized. Can have multiple
func init() {
	fmt.Println("calc/math.go ==> init() [1]")
}

func init() {
	fmt.Println("calc/math.go ==> init() [2]")
}

// Add - returns sum of two integers
func Add(numbers ...int) (int, error) {
	sum := 0
	if len(numbers) < 2 {
		return sum, errors.New("provide at least 2 numbers")
	}
	for _, num := range numbers {
		sum = sum + num
	}
	return sum, nil
}
