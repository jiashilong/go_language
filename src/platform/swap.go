package main

import (
	"fmt"
)

//值传递
func swap(x int32, y int32) {
	var t int32 = 0
	t = x
	x = y
	y = t
}

//多值返回
func swap2(x int32, y int32) (int32, int32) {
	return y, x
}

//引用传递
func swap3(x *int32, y *int32) {
	var t int32 = 0
	t = *x
	*x = *y
	*y = t
}

func main() {
	var x int32 = 10
	var y int32 = 20

	fmt.Println(x, y)
	swap(x, y)
	fmt.Println(x, y)

	fmt.Println(swap2(x, y))

	swap3(&x, &y)
	fmt.Println(x, y)
}
