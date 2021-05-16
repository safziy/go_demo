package main

import (
	"fmt"
	"go_demo/json_demo"
	"go_demo/redis_demo"
	"go_demo/struct_demo"
)

func main() {
	// 结构体的定义和初始化
	struct_demo.TestStructConstruct()
	fmt.Println("=====================分割线=====================")
	// 结构体嵌套和匿名成员
	struct_demo.TestStructNesting()
	fmt.Println("=====================分割线=====================")
	// JSON
	json_demo.TestJson()
	fmt.Println("=====================分割线=====================")
	// Redis
	redis_demo.TestRedisConnection()
	fmt.Println("=====================分割线=====================")
}
