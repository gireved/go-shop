package main

import (
	"go-shop/config"
	"go-shop/internal/startup"
	"go-shop/pkg/logger"
)

func main() {
	loading()
}

func loading() {
	config.InitConfig()
	startup.InitMySQL()
	startup.InitLogger()

	// 结束时同步日志
	defer logger.Sync()
}
