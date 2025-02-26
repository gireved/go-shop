package startup

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-shop/internal/routers"
	"go-shop/internal/service"
)

// NewGinEngine 创建 Gin 实例，注入中间件，并注册路由
func NewGinEngine(middlewares []gin.HandlerFunc, productService service.ProductService) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...) // 注入 Prometheus 监控中间件

	// 绑定 Product 相关路由
	routers.SetupRoutes(router, productService)

	// Prometheus 指标接口
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return router
}
