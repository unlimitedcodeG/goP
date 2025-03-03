package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)

	fmt.Printf("Splitted in slice : %v\n", sl)

	for _, val :=range sl {
		fmt.Printf("%s - ", val)
	}
}
