package main

import (
	"fmt"
)

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	//多条件判断
	a := 10
	b := 20

	if a >= 0 && b <= 20 {
		fmt.Println(b)
	}
}
