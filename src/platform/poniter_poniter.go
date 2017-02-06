package main

import "fmt"

func main() {
	var n int = 100
	var p1 *int
	var p2 **int

	p1 = &n
	p2 = &p1

	fmt.Printf("%d, %d, %d\n", n, *p1, **p2)
}
