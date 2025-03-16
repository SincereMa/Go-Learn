package main

import (
	"testing"
)

// 测试函数
func TestAdd(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{100, -200, -100},
	}

	// 遍历测试用例并执行测试
	for _, tt := range tests {
		// 执行测试
		actual := Add(tt.a, tt.b)
		// 比较结果
		if actual != tt.expected {
			t.Errorf("Add(%d, %d) = %d, expected %d", tt.a, tt.b, actual, tt.expected)
		}
	}
}

func BenchmarkLongRunningFunction(b *testing.B) {
	// b.N 是由基准测试框架设置的迭代次数
	for i := 0; i < b.N; i++ {
		LongRunningFunction()
	}
}
