package main



import "fmt"


func trim(s string) string{
	if len(s) ==0{
		return s
	}


	left, right := 0, len(s)-1;

	// 找到第一个非空格字符
	for left < right&& s[left] == ' '{
		left ++
	}

	// 找到最后一个非空格字符
	for right >=left && s[right] == ' '{
		right --
	}
	return s[left:right+1]
}


func main(){
	fmt.Println(trim("  hello world  "))
}