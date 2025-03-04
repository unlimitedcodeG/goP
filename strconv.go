package main

import (
	"fmt"
	"strconv"
)

func StrConv() {
	str := "42"

	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Print("转换失败")
	} else {
		fmt.Println("转换成功", num)
	}
}
