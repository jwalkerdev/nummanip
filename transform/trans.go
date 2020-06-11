package trans

import "fmt"

// init - runs when the package is initialized. Can have multiple
func init() {
	fmt.Println("transform/trans.go ==> init() [1]")
}

// Trans1 - returns sum of two integers
func Trans1(i int) int {
	return i * i
}
