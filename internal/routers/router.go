package routers

import (
	"github.com/gin-gonic/gin"
	"go-shop/internal/routers/routes"
	"go-shop/internal/service"
)

func NewRouter(productService service.ProductService) *gin.Engine {
	router := gin.Default()
	routes.SetupRoutes(router, productService)
	return router
}
