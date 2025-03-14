package api

import (
	"multimod/utils" // 导入 utils 包
)

// Greet 公开的 API 函数
// 使用 utils 包中的 ReverseString 函数
func Greet(name string) string {
	reversedName := utils.ReverseString(name)
	return "Hello, " + reversedName + "!"
}
