package main

import "fmt"

func main() {
	var n int = 10
	var p *int

	fmt.Printf("%d\n", &n)
	fmt.Printf("%d\n", p)

	p = &n
	fmt.Printf("%d, %d\n", p, *p)
}
