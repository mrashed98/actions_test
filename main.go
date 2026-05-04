// Go Project Skeleton.
package main

import (
	"fmt"
)

func SumOfTwoNumber(a int, b int) int {
	// SumOfTwoNumber of two numbers a and b
	// Example: a=1, b=2
	// Returns 3 i.e 1+2
	return a + b
}

func main() {
	// Entry Point just prinitng the value return from SumOfTwoNumber

	fmt.Println(SumOfTwoNumber(1, 2))
}
