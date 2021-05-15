package main

import (
	"fmt"
	"unicode/utf8"
)

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func main() {
	fmt.Println("Hello Go!")
	p := new(int)
	fmt.Println(p)
	fmt.Println(*p)

	fmt.Println(byte(4))

	a := 9
	fmt.Printf("%d %[1]o %#[1]o %[1]x  %#[1]x\n", a)

	fmt.Println("======================")
	s := "Hello, 中国"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	fmt.Println("======================")

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\t%d\n", i, r, size)
		i += size
	}
	fmt.Println("======================")
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Println("======================")
	fmt.Printf("% x\n", s)
	r := []rune(s)
	fmt.Printf("%x\n", r)
	fmt.Println(string(1234567))
}
