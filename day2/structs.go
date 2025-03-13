package main

import "fmt"

// 定义一个名为 User 的结构体
type User struct {
	ID    int
	Name  string
	Email string
}

// 定义一个名为 UserEmail 的结构体 包含User
type UserEmail struct {
	User
	Email string
}

func main() {
	// 创建 User 结构体实例
	user1 := User{ID: 1, Name: "Alice", Email: "alice@example.com"}

	// 复制 user1 到 user2
	user2 := user1

	// 修改 user2 的 Name 字段
	user2.Name = "Bob"

	// 打印 user1 和 user2，观察值类型的特性
	fmt.Println("user1:", user1) // 输出: user1: {1 Alice alice@example.com}
	fmt.Println("user2:", user2) // 输出: user2: {1 Bob alice@example.com}

	// 创建 UserEmail 结构体实例
	user3 := UserEmail{User: user1, Email: "alice_email@example.com"}
	// 打印 user3
	fmt.Println("user3:", user3) //user3: {{1 Alice alice@example.com} alice_email@example.com}

}
