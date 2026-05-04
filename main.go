// Go Project Skeleton.
package main

import "fmt"

// SumOfTwoNumber returns the sum of two integers.
func SumOfTwoNumber(a int, b int) int {
	return a + b
}

// SubtractTwoNumbers returns the difference of two integers.
func SubtractTwoNumbers(a int, b int) int {
	return a - b
}

// MultiplyTwoNumbers returns the product of two integers.
func MultiplyTwoNumbers(a int, b int) int {
	return a * b
}

func main() {
	fmt.Printf("%d\n", SumOfTwoNumber(1, 2))
}
