package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	if a == 10 {
		if b <= 20 {
			fmt.Println("small")
		} else {
			fmt.Println("big")
		}
	} else {
		fmt.Println("hehe")
	}
}
