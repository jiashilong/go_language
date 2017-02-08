package main

import "fmt"

func f(s string) {
	for i := 1; i <= 10; i++ {
		fmt.Println(s, ":", i)
	}
}

func f2() {
	for true {
		fmt.Println("f2")
	}
}

func f3() {
	for true {
		fmt.Println("f3")
	}
}

func main() {
	f("direct")

	go f2()

	go func() {
		f3()
	}()

	var input string = ""
	fmt.Scanln(&input)
	fmt.Println("done")
}
