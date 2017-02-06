package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func swap(a int, b string) (string, int) {
	return b, a
}

func main() {
	fmt.Println(max(1, 2))
	fmt.Println(min(1, 2))

	var a, b = swap(100, "hello")
	fmt.Printf("a=%s, b=%d\n", a, b)
}
