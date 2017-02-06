package main

import (
	"fmt"
	"strings"
)

func main() {
	var greeting string = "Hello World"

	fmt.Printf("%s\n", greeting)

	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%c ", greeting[i])
	}
	fmt.Println()

	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%x ", greeting[i])
	}
	fmt.Println()

	var greetings = []string{"Hello", "World"}
	fmt.Println(strings.Join(greetings, ","))
}
