package main

import (
	"fmt"
)

func main() {
	var i int    // 默认为 0
	var s string // 默认为 ""
	var b bool   // 默认为 false
	var p *int   // 默认为 nil

	fmt.Println(i, s, b, p)
	// 可以在声明变量时，指定变量类型和初始值。
	var number int = 10
	fmt.Println(number) // 输出: 10
}
