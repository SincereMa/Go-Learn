package main

import "fmt"

// 定义一个名为 Stringer 的接口
type Stringer interface {
	String() string
}

// 定义一个名为 Stringer2 的接口 嵌套Stringer
type Stringer2 interface {
	Stringer
	String2() string
}

type User struct {
	ID    int
	Name  string
	Email string
}

// User 类型实现了 Stringer 接口
func (u User) String() string {
	return fmt.Sprintf("User ID: %d, Name: %s", u.ID, u.Name)
}

// PrintStringer 函数接受任何实现了 Stringer 接口的类型
func PrintStringer(s Stringer) {
	fmt.Println(s.String())
}

// PrintStringer2 函数接受任何实现了 Stringer2 接口的类型
func PrintStringer2(s Stringer2) {
	fmt.Println(s.String2())
}

func main() {
	user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}

	// User 类型隐式地实现了 Stringer 接口，可以传递给 PrintStringer 函数
	PrintStringer(user) // 输出: User ID: 1, Name: Alice

	// User 类型没有实现 Stringer2 接口，因此不可以传递给 PrintStringer2 函数
	// PrintStringer2(user)
}
