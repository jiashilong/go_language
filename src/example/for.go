package main

import (
	"fmt"
)

func main() {
	i := 1 //声明并赋值，类型推导
	//单条件循环
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	//多条件循环
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}

}
