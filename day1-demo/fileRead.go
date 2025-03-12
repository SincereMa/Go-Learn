package main

import (
	"fmt"
	"os"
)

// readFileContent 函数读取指定路径的文件内容
// 参数：filePath 文件路径
// 返回值：文件内容（字符串）和错误信息（如果读取失败）
func readFileContent(filePath string) (string, error) {
	data, err := os.ReadFile(filePath) // 使用 os.ReadFile 读取文件
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err) // 如果出错，返回空字符串和格式化错误
	}
	return string(data), nil // 将字节切片转换为字符串，并返回 nil 错误（表示成功）
}

func main() {
	// 调用 readFileContent 函数，并获取文件内容和错误信息
	content, err := readFileContent("exist.txt")
	if err != nil {
		fmt.Println(err) // 如果有错误，直接打印错误信息
		return           //并退出程序
	}
	fmt.Println("File content:", content) // 如果没有错误，打印文件内容

	content, err = readFileContent("example.txt")
	if err != nil {
		fmt.Println(err) // 如果有错误，直接打印错误信息
		return           //并退出程序
	}
}
