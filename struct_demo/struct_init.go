package struct_demo

import "fmt"

type Animal struct {
	// 属性
	ID   int
	Name string
	Age  int
}

func (a *Animal) Eat() {
	// 方法
	fmt.Println("Animal Eat : yummy  yummy, ID : ", a.ID)
}

// TestStructConstruct 结构体的初始化方法
func TestStructConstruct() {
	// 初始方法1
	var animal1 Animal
	animal1.ID = 1
	animal1.Name = "kiki"
	animal1.Age = 10
	fmt.Println("init type1: ", animal1)
	animal1.Eat()

	// 初始方法2
	animal2 := Animal{ID: 2, Name: "kitty", Age: 13}
	fmt.Println("init type2: ", animal2)
	animal2.Eat()

	// 初始化方法3 new方法创建
	animal3 := new(Animal)
	animal3.ID = 3
	animal3.Name = "tom"
	animal3.Age = 17
	fmt.Println("init type3: ", animal3)
	animal3.Eat()
}
