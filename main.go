package main

import "fmt"

var (
	myRes = func(a int, b int) int {
		return a - b
	}
)

func main() {
	res1 := func(a int, b int) int {
		return a + b
	}(10, 25)

	fmt.Printf("res1=%d\n", res1)

	// 匿名函数 付给变量用变量来调用  可多次调用 可是作用地域有限

	res2 := func(a int, b int) int {
		return a * b
	}
	res3 := res2(10, 25)
	fmt.Printf("res3=%d\n", res3)

	res4 := myRes(10, 25)
	fmt.Printf("res4=%d\n", res4)
}
