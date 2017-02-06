package main

import "fmt"
import "unsafe"

func main() {
	const LENGTH int32 = 10
	const WIDTH int32 = 20

	fmt.Printf("%d, %d, %d\n", LENGTH, WIDTH, LENGTH*WIDTH)
	fmt.Printf("%T, %d\n", WIDTH, unsafe.Sizeof(LENGTH*WIDTH))
}
