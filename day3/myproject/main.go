package main // 主包声明

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux" // 导入gorilla/mux路由包

	"day3/myproject/internal" // 导入内部包
)


func main() {

	// 调用内部函数addInts并打印结果
	fmt.Println(internal.AddInts(1, 2))
	
	// 创建一个新的路由器实例
	r := mux.NewRouter()

	// 注册根路径的处理函数
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 向响应写入"Hello, world!"
		fmt.Fprintf(w, "Hello, world!")
	})

	// 启动HTTP服务器，监听8080端口
	http.ListenAndServe(":8080", r)

}
