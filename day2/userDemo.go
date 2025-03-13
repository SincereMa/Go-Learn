package main

import (
	"errors"
	"fmt"
)

// User 结构体定义
type User struct {
	ID    int
	Name  string
	Email string
}

// 定义错误信息
var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidUser  = errors.New("invalid user data")
)

// users  slice 切片存储所有用户（模拟数据库）
var users []User

// nextID 用于生成唯一的自增用户 ID
var nextID = 1

// CreateUser 创建新用户（指针接收者，修改 users 列表）
func (u *User) CreateUser() error {
	// 简单验证
	if u.Name == "" || u.Email == "" {
		return ErrInvalidUser
	}

	u.ID = nextID
	nextID++
	users = append(users, *u) // 注意：这里需要存储 u 的副本，而不是指针
	return nil
}

// GetUserByID 根据 ID 获取用户（值接收者，不修改数据）
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return User{}, ErrUserNotFound
}

// UpdateUser 更新用户信息（指针接收者，修改 users 列表中的用户）
func (u *User) UpdateUser() error {
	if u.Name == "" || u.Email == "" {
		return ErrInvalidUser
	}

	for i, existingUser := range users {
		if existingUser.ID == u.ID {
			users[i] = *u // 更新用户信息
			return nil
		}
	}
	return ErrUserNotFound
}

// DeleteUserByID 根据 ID 删除用户（修改 users 列表）
func DeleteUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			// 从 users 切片中移除元素
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return ErrUserNotFound
}

func main() {
	// 创建用户
	user1 := User{Name: "Alice", Email: "alice@example.com"}
	err := user1.CreateUser()
	if err != nil {
		fmt.Println("Error creating user:", err)
	}

	user2 := User{Name: "Bob", Email: "bob@example.com"}
	err = user2.CreateUser()
	if err != nil {
		fmt.Println("Error creating user:", err)
	}

	// 获取用户
	retrievedUser, err := GetUserByID(user1.ID)
	if err != nil {
		fmt.Println("Error getting user:", err)
	} else {
		fmt.Println("Retrieved user:", retrievedUser)
	}

	// 更新用户
	user1.Name = "Alice Updated"
	err = user1.UpdateUser()
	if err != nil {
		fmt.Println("Error updating user:", err)
	}

	// 再次获取用户，验证更新
	retrievedUser, err = GetUserByID(user1.ID)
	if err != nil {
		fmt.Println("Error getting user:", err)
	} else {
		fmt.Println("Retrieved user after update:", retrievedUser)
	}

	// 删除用户
	err = DeleteUserByID(user2.ID)
	if err != nil {
		fmt.Println("Error deleting user:", err)
	}

	// 尝试获取已删除的用户
	_, err = GetUserByID(user2.ID)
	if err == ErrUserNotFound {
		fmt.Println("User successfully deleted")
	}
}