package main

import "fmt"

type Book struct {
	id      int
	title   string
	author  string
	subject string
}

func show(b Book) {
	fmt.Printf("%d, %s, %s, %s\n", b.id, b.title, b.author, b.subject)
}

//指针也是使用.访问成员变量
func show2(b *Book) {
	fmt.Printf("%d, %s, %s, %s\n", b.id, b.title, b.author, b.subject)
}

func main() {
	var b1 Book

	/* book 1 specification */
	b1.title = "Go Programming"
	b1.author = "Mahesh Kumar"
	b1.subject = "Go Programming Tutorial"
	b1.id = 6495407

	fmt.Printf("%d, %s, %s, %s\n", b1.id, b1.title, b1.author, b1.subject)
	show(b1)
	show2(&b1)
}
