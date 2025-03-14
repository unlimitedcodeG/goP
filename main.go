package main

import "fmt"

// var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	map1 := make(map[int]float32)

	map1[1] = 1.0

	for key, value := range map1 {
		// print key && value
		fmt.Printf("key is :%d - value is :%f\n", key, value)
	}

	for key := range map1 {
		fmt.Printf("key is %d\n", key)
	}

	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}

}
