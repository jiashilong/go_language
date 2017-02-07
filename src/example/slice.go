package main

import (
	"fmt"
)

func main() {
	s := make([]int, 3)
	fmt.Println("emp: ", s)

	for i := 0; i < len(s); i++ {
		s[i] = i + 1
	}
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	s = append(s, 4)
	s = append(s, 5, 6)
	fmt.Println("append: ", s)

	c := make([]int, len(s))
	copy(c, s)
	fmt.Println("copy: ", c)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	k := s[:5]
	fmt.Println("sl2 ", k)

	m := s[2:]
	fmt.Println("sl3 ", m)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		len := i + 1
		twoD[i] = make([]int, len)
		for j := 0; j < len; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
