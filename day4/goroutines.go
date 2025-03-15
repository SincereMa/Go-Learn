package main

import (
	"fmt"
	"time"
)

func myFunc() {
	fmt.Println("Hello from a Goroutine!")
}

func main() {
	go myFunc()             // 启动一个新的 Goroutine，执行 myFunc
	time.Sleep(time.Second) // 等待 Goroutine 执行完毕
	fmt.Println("Hello from the main function!")
}
