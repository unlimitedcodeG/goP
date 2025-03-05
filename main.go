package main

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func main() {
	println(DeferFunc2(1))
}
