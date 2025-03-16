package main

import (
	"fmt"
	"net/http"
)

// 假设这是一个需要优化的函数
// 它计算了斐波那契数列的第 n 项
// func fibonacci(n int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	// 故意使用低效的递归实现来掩饰性能问题
// 	return fibonacci(n-1) + fibonacci(n-2)
// }

func fibonacci(n int) int {
	memo := make(map[int]int)
	return fibonacciMemo(n, memo)
}
func fibonacciMemo(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if _, ok := memo[n]; !ok {
		memo[n] = fibonacciMemo(n-1, memo) + fibonacciMemo(n-2, memo)
	}
	return memo[n]
}

// HTTP handler，使用 fibonacci 函数
func fibHandler(w http.ResponseWriter, r *http.Request) {
	// 获取查询参数 n
	nStr := r.URL.Query().Get("n")
	var n int
	fmt.Scan(nStr, &n) // 简单起见，不处理错误

	// 计算斐波那契数列
	result := fibonacci(n)

	// 将结果写入响应
	fmt.Fprintf(w, "Fibonacci(%d) = %d\n", n, result)
}

func main() {
	// 注册 HTTP handler
	http.HandleFunc("/fib", fibHandler)

	// 启动 HTTP 服务器（同时启用 pprof 端点）
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
