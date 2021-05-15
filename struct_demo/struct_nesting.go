package struct_demo

import "fmt"

// Cat 结构体嵌套
type Cat struct {
	Ani Animal
	Color string
}

// Dog 匿名成员
type Dog struct {
	Animal
	kind string
}

// TestStructNesting 结构体的嵌套
func TestStructNesting()  {
	cat := Cat{
		Ani: Animal{
			ID: 4,
			Name: "timmy",
			Age: 5,
		},
		Color: "red",
	}
	fmt.Println(cat, cat.Ani.ID)
	cat.Ani.Eat()

	dog := Dog{
		kind: "Garfield",
		Animal: Animal{
			ID: 6,
			Name: "pet",
			Age: 3,
		},
	}
	fmt.Println(dog, dog.ID)
	dog.Eat()
}
