package main

import (
	"fmt"
	"testing"
)

// 表格驱动测试 fibonacci 函数
func TestFibonacci(t *testing.T) {
	var tests = []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{10, 55},
		{20, 6765}, // 较大的数
	}

	for _, tt := range tests {
		actual := fibonacci(tt.n)
		if actual != tt.expected {
			t.Errorf("fibonacci(%d) = %d; expected %d", tt.n, actual, tt.expected)
		}
	}
}

// 基准测试 fibonacci 函数
func BenchmarkFibonacci(b *testing.B) {
	// 基准测试不同的 n 值
	for n := 10; n <= 30; n += 10 {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibonacci(n)
			}
		})
	}
}
