package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int8
}

func zeroValue(n int) {
	n = 0
}

func zeroPtr(n *int) {
	*n = 0
}

func main() {
	var n int = 10
	fmt.Println(n)

	zeroValue(n)
	fmt.Println(n)

	zeroPtr(&n)
	fmt.Println(n)

	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{name: "Alice", age: 30})
	fmt.Println(Person{name: "Jarry"})
	fmt.Println(Person{age: 100})

	fmt.Println(&Person{name: "Ann", age: 40})
	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)

}
