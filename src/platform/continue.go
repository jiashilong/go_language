package main

import (
	"fmt"
)

func main() {
	const NUM int = 20
	var i int = 10

	for i <= NUM {
		if i == 15 {
			i++
			continue
		}

		fmt.Println(i)
		i++
	}
}
