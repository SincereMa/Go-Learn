package main

import (
	"fmt"
)

func main() {
	// 无缓冲 Channel
	ch := make(chan int)
	
	go func()  {
		ch <- 10  // 发送数据
	}()

	value := <-	ch  // 接收数据
	fmt.Println(value) // 输出 10

	// 有缓冲 Channel
	chBuffered := make(chan string, 2)

	chBuffered <- "Hello"
	chBuffered <- "World"

	fmt.Println(<-chBuffered)	// 输出 Hello
	fmt.Println(<-chBuffered)	// 输出 World

	close(ch)
	close(chBuffered)
}