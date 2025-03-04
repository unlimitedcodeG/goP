package main

import "os"

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

func main() {
	a := "G"
	print(a)
	print(HOME)
	print("hhhh\n")
	print(USER)
	print("hhhh\n")
	print(GOROOT)
	f1()
}
