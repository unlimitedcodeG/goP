package main

import (
	"fmt"
	"sync"
)

func f(from string, wg *sync.WaitGroup) {
	defer wg.Done() // 任务完成时减少计数
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // 等待两个 Goroutine 完成

	go f("goroutine", &wg)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}("going")

	wg.Wait() // 等待所有 Goroutine 结束
	fmt.Println("done")
}
