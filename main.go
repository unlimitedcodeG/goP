package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) AreaPointer() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	var c Circle
	c.radius = 10.00

	// 调用圆形的Area方法
	fmt.Println("Area of circle:", c.Area())
	// 调用圆形的AreaPointer方法
	fmt.Println("Area of circle (pointer):", c.AreaPointer())
}
