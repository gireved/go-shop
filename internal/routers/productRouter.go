package routers

import (
	"github.com/gin-gonic/gin"
	"go-shop/internal/api"
	"go-shop/internal/service"
)

func NewProductRouter(svc service.ProductService) *gin.Engine {
	router := gin.Default()
	productHandler := api.NewProductHandler(svc)
	productGroup := router.Group("/product")
	{
		//router.GET("/:id", GetProduct)
		productGroup.POST("/", productHandler.CreateProductHandler())
		//router.PUT("/:id", UpdateProduct)
		//router.DELETE("/:id", DeleteProduct)
	}
	return router
}
