package main

import "fmt"

func main() {
	x := 10      // 自动推导类型为 int
	name := "Go" // 自动推导为 string 类型

	fmt.Printf("x的类型: %T, name的类型: %T\n", x, name)
}
