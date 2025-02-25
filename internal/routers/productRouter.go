package routers

import (
	"github.com/gin-gonic/gin"
	"go-shop/internal/api"
	"go-shop/internal/service"
)

/*type ProductRouter struct {
	svc            service.ProductService
	productHandler *api.ProductHandler
}

func NewProductRouter(svc service.ProductService) *ProductRouter {
	handler := api.NewProductHandler(svc)
	return &ProductRouter{
		svc:            svc,
		productHandler: handler,
	}
}

func (p *ProductRouter) SetupRouter() *gin.Engine {
	router := gin.Default()
	productGroup := router.Group("/product")
	{
		//router.GET("/:id", GetProduct)
		productGroup.POST("/", p.productHandler.CreateProductHandler())
		//router.PUT("/:id", UpdateProduct)
		//router.DELETE("/:id", DeleteProduct)
	}
	return router
}*/

// NewProductRouter 配置路由并返回 *gin.Engine
func NewProductRouter(svc service.ProductService) *gin.Engine {
	// 创建 Gin 引擎实例
	router := gin.Default()

	// 配置路由
	productHandler := api.NewProductHandler(svc)
	productGroup := router.Group("/product")
	{
		productGroup.POST("/", productHandler.CreateProductHandler())
	}

	return router
}
