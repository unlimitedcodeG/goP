package main

import "fmt"

// 这里 ch 是一个缓冲通道，容量为 3。即使没有接收者，前 3 次 ch <- 依然可以执行，不会阻塞，直到缓冲区满后才会阻塞。
func main() {
	ch := make(chan int, 3) // 创建一个带 3 个容量的通道

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(<-ch) // 读取第一个元素
	fmt.Println(<-ch) // 读取第二个元素
	fmt.Println(<-ch) // 读取第三个元素
}
