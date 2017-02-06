package main

import (
	"fmt"
)

func main() {
	var i, j, k int
	var c, ch byte
	var f, salary float32
	var d int = 42

	var x, y, z = 1, false, "hello"

	var m = 100
	n := 200
	k = 300

	str := `hello
    world
v2.0`

	fmt.Printf("%d, %d, %d\n", i, j, k)
	fmt.Printf("%d, %d\n", c, ch)
	fmt.Printf("%f, %f\n", f, salary)
	fmt.Printf("d=%d, d is type of %T\n", d, d)

	fmt.Println(x, y, z)
	fmt.Printf("%T, %T, %T\n", x, y, z)

	fmt.Println(m, n, k)
	fmt.Printf("%T, %T, %T\n", m, n, k)

	fmt.Println(str)
}
