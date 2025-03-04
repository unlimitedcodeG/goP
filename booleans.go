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

	var first int = 10
	var cond int

	if first <= 0 {
		fmt.Printf("first is less than or equal to 0\n")
	}else if first>0&&first<5{
		fmt.Printf("first is between 0 and 5 \n")
	}else {
		fmt.Printf("first is 5 or greater\n")
	}
	if cond =5;cond>10{
		fmt.Printf("cond is g than 10\n")
	}else {
		fmt.Printf("cond is not greater than 10 ")
	}
}
