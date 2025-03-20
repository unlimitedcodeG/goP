package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()")
}

type SuperMan struct {
	Human
	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan eat ")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}
func main() {
	h := Human{"zhang3", "female"}

	fmt.Println(h)

	sh := SuperMan{Human{"superman", "male"}, 1}

	fmt.Println(sh)
	sh.Eat()
	h.Eat()
	sh.Fly()
	sh.Walk()
}
