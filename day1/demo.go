package main

import (
	"fmt"
)

func main() {
	fmt.Println("查看变量的自动类型推导")
	PlayVars()

	fmt.Println("查看变量的自动类型推导")
	PlayTypes()

	fmt.Println("演示多返回值")
	fmt.Println(divide(10, 2)) // 输出: 5, nil

	fmt.Println("演示多返回值和 error")
	fmt.Println(divide(10, 0)) // 输出: 0, cannot divide by zero; 需要显示检查和处理错误
}

func PlayVars() {
	x := 10      // 自动推导类型为 int
	name := "Go" // 自动推导为 string 类型

	fmt.Printf("x的类型: %T, name的类型: %T\n", x, name)
}

func PlayTypes() {
	var i int    // 默认为 0
	var s string // 默认为 ""
	var b bool   // 默认为 false
	var p *int   // 默认为 nil

	fmt.Println(i, s, b, p)
	// 可以在声明变量时，指定变量类型和初始值。
	var number int = 10
	fmt.Println(number) // 输出: 10
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero") // 创建一个错误
	}
	return a / b, nil // 无错误时返回 nil
}
