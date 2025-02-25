package routes

import (
	"github.com/gin-gonic/gin"
	"go-shop/internal/service"
)

func SetupRoutes(router *gin.Engine, productService service.ProductService) {

	// 加载产品路由
	SetupProductRoutes(router, productService)

}
