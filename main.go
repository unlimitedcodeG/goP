package main

import "fmt"

// declare type func

type cb func(int) int

func main() {
	testCallBack(1, callBack) // testCallBack

}

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("I am callback x:%d", x)
	return x
}
