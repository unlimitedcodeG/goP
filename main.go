package main

import "fmt"

func main() {

	var s []string

	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	s = make([]string, 3)

	fmt.Println("emp:", s, "len:", "cap", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len: ", len(s))
	// s = make([]string, 3)
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
}
