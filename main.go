package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

type Person struct {
	Id   int32
	Name string
	Tags []string
}

func main() {
	// 创建对象
	p := &Person{
		Id:   123,
		Name: "Alice",
		Tags: []string{"gamer", "developer"},
	}

	// 序列化（编码）
	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal("序列化失败:", err)
	}

	fmt.Println("Protobuf 序列化数据:", data)

	// 反序列化（解码）
	var newPerson Person
	err = proto.Unmarshal(data, &newPerson)
	if err != nil {
		log.Fatal("反序列化失败:", err)
	}

	fmt.Println("解析后的数据:", newPerson)
}
