package main

import (
	"fmt"
)

func main() {
	var a string = "hello world"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	f := "short" //声明并赋值变量，类型推导
	fmt.Println(f)
}
