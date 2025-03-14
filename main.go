package main

import "fmt"

// var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}

	fmt.Println(m)

	v1 := m["apple"]
	v2, ok := m["pear"]

	m["apple"] = 5

	println(v1)
	println(v2)
	println(ok)
}
