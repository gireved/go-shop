package routers

import (
	"github.com/gin-gonic/gin"
	"go-shop/internal/api"
	"go-shop/internal/service"
)

// SetupProductRoutes 绑定产品相关路由
func SetupProductRoutes(router *gin.Engine, svc service.ProductService) {
	productHandler := api.NewProductHandler(svc)

	// 创建 /product 分组
	productGroup := router.Group("/product")
	{
		productGroup.POST("/", productHandler.CreateProductHandler())
		productGroup.PUT("/", productHandler.UpdateProductHandler())
		productGroup.DELETE("/", productHandler.DeleteProductHandler())
	}
}
