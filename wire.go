//go:build wireinject

package go_shop

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go-shop/internal/repository/dao"
	"go-shop/internal/service"
	"go-shop/internal/startup"
	g "go-shop/pkg/ginx"
)

// InitializeProductRouter 初始化 Product 相关的 Gin 路由
func InitializeProductRouter() *gin.Engine {
	wire.Build(
		startup.GetDB,              // 获取数据库实例
		dao.NewProductDao,          // 注入 Product Dao
		service.NewProductService,  // 注入 Product Service
		startup.InitGinMiddlewares, // 注入 Gin 中间件
		g.NewGinEngine,             // 创建 Gin 引擎
	)
	return nil
}
