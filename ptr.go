package main

import (
	"fmt"
)

// const i = 5

// ptr := &i //error: cannot take the address of i
// ptr2 :=&10
func PtrMain() {
	// 如果你真的需要字符串或常量的地址，你需要存到变量里：
	// ptr := &"hello"  //error: cannot take the address of "hello"
	s := "good bye" // 使用变量
	var p *string = &s
	*p = "ciao"
	fmt.Printf("*p:%v\n", *p)
	fmt.Printf("s:%v\n", s)
	fmt.Printf("p:%v\n", p)
	// const x = 42
	// ptr := &x // ❌ 编译错误：cannot take the address of x
	var num = 42
	ptr := &num
	fmt.Printf("ptr:%v\n", ptr)
	// const i = 5
	// fmt.Printf(&i) //invalid operation: cannot take address of i (untyped int constant 5)

}
