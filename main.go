package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	Level int
}
// 如果说类的属性首字母大写，表示该属性是对外能够访问的，否则的话 只能够🥱类内部访问
func (this Hero) show() {
	fmt.Println("Name =", this.Name)
	fmt.Println("Ad =", this.Ad)
	fmt.Println("Level =", this.Level)
}

func (this Hero) GetName() string {
	return this.Name
}

func (this Hero) SetName(newName string) {
	this.Name = newName
	// return newName
}

func main() {
	hero := Hero{Name: "zhang3", Ad: 100, Level: 1}

	hero.SetName("lisi")
	hero.show()
}
