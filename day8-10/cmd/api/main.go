package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	// 初始化 Viper
	viper.SetConfigName("config")    //配置文件名（不带扩展名）
	viper.SetConfigType("yaml")      //配置文件类型
	viper.AddConfigPath("./configs") //配置文件路径
	err := viper.ReadInConfig()      //读取配置文件
	if err != nil {
		panic(err)
	}

	userHandler, err := InitializeUserHandler() // 使用 Wire 生成的函数
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
