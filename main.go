package main

import "fmt"

func main() {
	var value = [3][2]int{{1, 2}, {3, 4}, {5, 6}}

	for i, v := range value {
		for j, v2 := range v {
			fmt.Printf("value[%v][%v]=%v\t", i, j, v2)
		}
		fmt.Print(v)
		fmt.Println()
	}
}
