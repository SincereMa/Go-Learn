package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// FileData 结构体，用于存储文件名和内容
type FileData struct {
	Name    string
	Content string
	Err     error
}

// processFile 函数处理单个文件， 并将结果发送到 Channel
func processFile(filePath string, resultChan chan<- FileData) { // 使用单向 Channel，限制只能发送
	content, err := os.ReadFile(filePath)
	resultChan <- FileData{Name: filepath.Base(filePath), Content: string(content), Err: err}
}

func main() {
	dir := "./test_files" // 假设要处理的文件都在这个目录下

	// 创建测试文件
	err := createTestFiles(dir)
	if err != nil {
		fmt.Println("创建测试文件出错", err)
		os.Exit(1)
	}

	// 创建一个有缓冲的 Channel
	resultChan := make(chan FileData, 10) // 缓冲区大小可以根据实际情况调整
	var wg sync.WaitGroup                 // 用于等待所有 Goroutine 完成

	// 遍历目录， 为每个文件启动一个 Goroutine
	files, err := os.ReadDir(dir) // 使用 os.ReadDir 读取目录下的文件
	if err != nil {
		fmt.Println("读取目录失败", err)
		os.Exit(1)
	}
	for _, file := range files {
		if !file.IsDir() { // 忽略子目录
			filePath := filepath.Join(dir, file.Name())
			wg.Add(1) // 增加 WaitGroup 计数器
			go func(fp string) {
				defer wg.Done()             // Goroutine 完成时减少计数器
				processFile(fp, resultChan) // 处理文件
			}(filePath)
		}
	}
	// 启动一个 Goroutine 来关闭 Channel
	go func() {
		wg.Wait()         // 等待所有文件处理 Goroutine 完成
		close(resultChan) // 关闭 Channel
	}()

	// 使用 select 监听 resultChan 和超时
	timeout := time.After(5 * time.Second) // 设置超时时间
	for {
		select {
		case result, ok := <-resultChan: // 从结果通道接收文件数据
			if !ok {
				// Channel 已关闭，所有文件处理完成
				fmt.Println("所有文件处理完成！")
				return
			}
			if result.Err != nil {
				fmt.Printf("处理文件 %s 出错: %v\n", result.Name, result.Err)
			} else {
				fmt.Println("文件内容读取成功", result.Name)
			}
		case <-timeout:
			fmt.Println("处理文件操作超时")
			return
		}
	}
}

// 创建测试文件
func createTestFiles(dir string) error {
	// 确保目录存在
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			panic(err)
		}
	}

	for i := 1; i <= 3; i++ {
		fileName := fmt.Sprintf("file%d.txt", i)
		content := []byte(fmt.Sprintf("This is the content of %s", fileName))
		err := os.WriteFile(filepath.Join(dir, fileName), content, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
