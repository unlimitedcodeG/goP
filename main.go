package main

import "fmt"

// 迭代器的支持
// 在 Go 1.23 版本之前，range 语句通常只能用于切片、数组、映射（map）、字符串等类型。Go 1.23 版本的更新引入了更通用的迭代器支持，这意味着我们可以使用 range 语法来遍历更加灵活和广泛的数据结构。

// 具体变化
// range 能够遍历更多类型：不仅仅是数组、切片、映射和字符串，还可以用来遍历其他可以实现迭代器接口的自定义数据类型。
// 自定义迭代器：我们可以定义自己实现迭代器的类型，并用 range 来遍历这些自定义的迭代器。

type MyIterator struct {
	items []string
	index int
}

// 实现 Next 方法，返回当前元素和是否有下一个元素
func (it *MyIterator) Next() (string, bool) {
	if it.index < len(it.items) {
		item := it.items[it.index]
		it.index++
		return item, true
	}
	return "", false
}

func main() {
	// 创建自定义迭代器
	iter := &MyIterator{items: []string{"apple", "banana", "cherry"}}

	// 使用range
	for item, ok := iter.Next(); ok; item, ok = iter.Next() {
		fmt.Println(item)
	}
}
