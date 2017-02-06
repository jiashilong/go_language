package main

import (
	"fmt"
)

func main() {
	const NUM int = 20
	var i int = 10

	for i <= NUM {
		fmt.Println(i)
		i++
		if i > 15 {
			break
		}
	}
}
