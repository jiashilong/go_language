package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 3)

	go func() {
		messages <- "100"
		messages <- "200"
		messages <- "300"
	}()

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	//fmt.Println(<-messages) //goroutines are asleep - deadlock
}
