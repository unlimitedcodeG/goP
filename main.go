package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPU 核心数:", runtime.NumCPU())
}
