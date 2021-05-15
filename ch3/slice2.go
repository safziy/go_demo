package main

import "fmt"

type Book struct {
	id   int32
	name string
}

func main() {
	s := make([]*Book, 0, 100)
	s = append(s, &Book{id: 1, name: "aaa"})
	s = append(s, &Book{id: 2, name: "bbb"})
	s = append(s, &Book{id: 3, name: "ccc"})
	s = append(s, &Book{id: 4, name: "ddd"})
	fmt.Println(s)
	s[2] = nil
	fmt.Println(s)
	s = append(s, &Book{id: 5, name: "eee"})
	fmt.Println(s)
}
