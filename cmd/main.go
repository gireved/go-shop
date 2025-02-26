package main

import (
	goshop "go-shop"
	"go-shop/internal/startup"
	"log"
)

func main() {
	startup.Init()
	// 使用 Wire 生成的注入器函数初始化路由器
	router := goshop.InitializeProductRouter()
	// 启动服务器
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}

}
