//go:build wireinject

package go_shop

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go-shop/internal/repository/dao"
	"go-shop/internal/routers"
	"go-shop/internal/service"
	"go-shop/internal/startup"
)

// 初始化依赖注入
func InitializeProductRouter() *gin.Engine {
	wire.Build(
		startup.GetDB,
		dao.NewProductDao,
		service.NewProductService,
		routers.NewRouter,
	)
	return nil
}
