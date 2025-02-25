package main

import (
	go_shop "go-shop"
	"go-shop/config"
	"go-shop/internal/startup"
	"go-shop/pkg/logger"
	"log"
)

func main() {
	loading()
	// 使用 Wire 生成的注入器函数初始化路由器
	router := go_shop.InitializeProductRouter()
	// 启动服务器
	if err := router.Run(); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}

}

func loading() {
	config.InitConfig()
	startup.InitMySQL()
	startup.InitLogger()

	// 结束时同步日志
	defer logger.Sync()
}
