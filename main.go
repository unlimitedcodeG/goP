package main

import "fmt"

func main() {

	var s []string

	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// s = make([]string, 3)

}
