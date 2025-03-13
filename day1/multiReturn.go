package main

import (
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero") // 创建一个错误
	}
	return a / b, nil // 无错误时返回 nil
}

func main() {

	fmt.Println("演示多返回值")
	fmt.Println(divide(10, 2)) // 输出: 5, nil

	fmt.Println("演示多返回值和 error")
	fmt.Println(divide(10, 0)) // 输出: 0, cannot divide by zero; 需要显示检查和处理错误
}
