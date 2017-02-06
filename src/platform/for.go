package main

import (
	"fmt"
)

func main() {
	const b int = 10

	// for a := 0; a < b; a++ {
	// 	fmt.Println(a)
	// }

	a := 0
	for a < b {
		fmt.Println(a)
		a++
	}
}
