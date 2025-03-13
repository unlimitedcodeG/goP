package main

import (
	"fmt"
)

func main() {
	add := func(a, b int) int {
		return a + b
	}

	result := add(3, 5)
	fmt.Println("3+5 = ", result)

	// inside exec anonymous func
	multiply := func(x, y int) int {
		return x * y
	}

	product := multiply(4, 6)
	fmt.Printf("4 * 6 =%d  ", product)

	calculate := func(operation func(int, int) int, x, y int) int {
		return operation(x, y)
	}

	sum := calculate(add, 2, 8)
	fmt.Println("2+8=", sum)

	difference := calculate(func(a, b int) int {
		return a - b
	}, 10, 4)
	fmt.Println("10 - 4 =  ", difference)
}
