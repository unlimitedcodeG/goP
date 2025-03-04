package main

import (
	"fmt"
	"runtime"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isGreater(x, y int) bool {
	return x > y
}
func BoolElse() {
	bool1 := true
	if runtime.GOOS == "windows" {
		bool1 = false
	}
	var max int = 6
	if val := 10; val > max {
		//do somthing
	}
	if bool1 {
		fmt.Printf("The value is true\n")
	} else {
		fmt.Printf("The value is false\n")
	}

	if isGreater(1, 3) {
		fmt.Print("g")
	} else {
		fmt.Print("l")
	}
}
