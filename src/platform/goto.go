package main

import (
	"fmt"
)

func main() {
	const NUM int = 20
	var i int = 10

LOOP:
	for i <= NUM {
		if i == 15 {
			i++
			goto LOOP
		}

		fmt.Println(i)
		i++
	}
}
