package models

// 定义 models 的接口
type UserRepository interface {
	GetUserByID(id int) (*User, error)
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
