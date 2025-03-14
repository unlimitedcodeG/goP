package main

import "fmt"

// var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	select {
	case ch <- 3:
		fmt.Println("成功发送 3")
	default:
		fmt.Println("ch 已经满了 跳过发送 3 ")
	}
	ch <- 3
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}

}
