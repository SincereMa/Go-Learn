package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

func (u User) Display() {
	fmt.Printf("ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
}

// 值接收者方法：不会修改原始结构体
func (u User) MockName() {
	u.Name = "Mock"
}

// 指针接收者方法：可以修改原始结构体
func (u *User) ChangeEmail(newEmail string) {
	u.Email = newEmail
}

func main() {
	user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}

	// 调用值接收者方法
	user.MockName()
	// 观察 Name 是否改变
	user.Display() // 输出: ID: 1, Name: Alice, Email: alice@example.com

	// 调用指针接收者方法
	user.ChangeEmail("new_alice@example.com")
	// 观察 Email 是否改变
	user.Display() // 输出: ID: 1, Name: Alice, Email: new_alice@example.com
}
