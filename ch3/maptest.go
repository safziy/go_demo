package main

import "fmt"

func test() (i int) {
	defer func() {
		i++
	}()
	return 1
}

func test2() int {
	i := 1
	defer func() {
		i++
	}()
	return i
}

func main() {
	m := make(map[string]string)
	m["aa"] = "bb"
	fmt.Println(m)

	a, b := m["bb"]
	fmt.Println(a, b)

	c, d := m["aa"]
	fmt.Println(c, d)

	fmt.Println(test(), test2())
}

func test3() {
	defer recover()
	panic("heelo")
}
