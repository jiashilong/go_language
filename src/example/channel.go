package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	//send
	go func() {
		for true {
			messages <- "ping"
		}
	}()

	//receive
	for true {
		msg := <-messages
		fmt.Println(msg)
	}

}
