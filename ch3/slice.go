package main

import "fmt"

func main() {
	arr := []string{"aa", "bb", "cc", "dd", "ee", "ff"}
	fmt.Println(arr, cap(arr), len(arr))

	a1 := arr[1:3]
	fmt.Println(a1, cap(a1), len(a1))
	a2 := a1[:5]
	fmt.Println(a2, cap(a2), len(a2))

	str := "abcdefg"
	fmt.Println(str[1:3])

	a := [...]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T \t %T \n", a, s)

	s = make([]int, 5, 20)
	fmt.Println(s, cap(s), len(s))
}
