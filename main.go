package main

import "fmt"

// var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, c := range "hello" {
		fmt.Printf("index:%d,char: %c\n", i, c)
	}
}
