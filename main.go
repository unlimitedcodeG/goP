package main

import "fmt"

func main() {
	var i1 = 5

	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

	if *(intP) == 5 {
		fmt.Println("The value is 5")
	}
}
