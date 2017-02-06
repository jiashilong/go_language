package main

import (
	"fmt"
)

func main() {
	var arr [10]int

	for i := 0; i < 10; i++ {
		arr[i] = 100 + i
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	var ts = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", ts[i][j])
		}
		fmt.Println()
	}

}
