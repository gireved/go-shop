package main

import (
	"go-shop/config"
	"go-shop/internal/repository/dao"
	"go-shop/internal/routers"
	"go-shop/internal/service"
	"go-shop/internal/startup"
	"go-shop/pkg/logger"
)

func main() {
	loading()
	db := startup.GetDB()
	productDao := dao.NewProductDao(db)
	productService := service.NewProductService(productDao)
	r := routers.NewProductRouter(productService)

	r.Run()
}

func loading() {
	config.InitConfig()
	startup.InitMySQL()
	startup.InitLogger()

	// 结束时同步日志
	defer logger.Sync()
}
