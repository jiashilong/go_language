package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println("sum: ", sum)

	for i, v := range nums {
		fmt.Println(i, v)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	fmt.Println(kvs)

	for k, v := range kvs {
		fmt.Println(k, v)
	}

	for k, _ := range kvs {
		fmt.Println(k)
	}

	for _, v := range kvs {
		fmt.Println(v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
