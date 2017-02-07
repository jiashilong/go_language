package main

import (
	"fmt"
	"math"
)

const s string = "const"

func main() {
	fmt.Println(s)

	const n = 500000000
	fmt.Println(n)

	const d = 3e20 / n
	fmt.Printf("%f\n", d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
