package main

import "fmt"

// var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	nums := []int{2, 3, 4}

	for _, num := range nums {
		fmt.Println("value:", num)
	}

	for i := range nums {
		fmt.Println("index:", i)
	}

}
