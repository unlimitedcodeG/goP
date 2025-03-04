package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
)

var Pi float64
var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}
func trim(s string) string {
	if len(s) == 0 {
		return s
	}

	left, right := 0, len(s)-1

	// 找到第一个非空格字符
	for left < right && s[left] == ' ' {
		left++
	}

	// 找到最后一个非空格字符
	for right >= left && s[right] == ' ' {
		right--
	}
	return s[left : right+1]
}

var a string

// func n() { print(a) }

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() { print(a) }

//	func m() {
//		a := "G"
//		print(a)
//		n()
//	}
var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

func init() {
	Pi = 4 * math.Atan(1)
}

// 变量除了可以在全局声明中初始化，也可以在 init 函数中初始化。这是一类非常特殊的函数，它不能够被人为调用，而是在每个包完成初始化后自动执行，并且执行优先级比 main 函数高。
// 吐槽 竟然不能调用
func main() {
	a := "G"
	print(a)
	print(HOME)
	print("hhhh\n")
	print(USER)
	print("hhhh\n")
	print(GOROOT)
	f1()
	PtrMain()
	BoolElse()
}
