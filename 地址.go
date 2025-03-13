package main

import "fmt"

func address() {
	a := 1005
	b := &a
	fmt.Printf("%T %[1]v,%p\n", &a, b)

	c := *b

	fmt.Printf("%T %[1]v\n", c)
	fmt.Println(a == c, &a, &c)

	var p1 *int // save int type value memory address in nil is very danger

	fmt.Printf("%T %[1]p %[1]v\n", p1)

	p1 = &a
	fmt.Printf("%T %[1]p %[1]v\n", p1)

	fmt.Println(*p1)

	*p1 = 2000
	fmt.Println(a, c) // a= 2000 c = 1005

}
