package main

import "fmt"

const SIZE int = 5

func main() {
	var data = []int{1, 2, 3, 4, 5}
	var ptr [SIZE]*int

	for i := 0; i < SIZE; i++ {
		ptr[i] = &data[i]
	}

	for i := 0; i < SIZE; i++ {
		fmt.Println(*ptr[i])
	}
}
