package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.2345

	fmt.Println("old value of pointer:", num)

	pointer := reflect.ValueOf(&num)

	newValue := pointer.Elem()
	fmt.Println(newValue)
}
