package startup

import (
	"github.com/gin-gonic/gin"
	prometheus "go-shop/pkg/ginx/middleware"
)

func InitGinMiddlewares() []gin.HandlerFunc {
	pb := &prometheus.Builder{
		Namespace: "go_shop_product",
		Subsystem: "go_shop",
		Name:      "gin_http",
		Help:      "统计 GIN 的HTTP接口数据",
	}

	return []gin.HandlerFunc{
		pb.BuildResponseTime(),
		pb.BuildActiveRequest(),
	}
}
