package api

import (
	"github.com/gin-gonic/gin"
	"go-shop/internal/service"
	"go-shop/internal/types"
	"go-shop/pkg/logger"

	"net/http"
)

type ProductHandler struct {
	svc service.ProductService
}

func NewProductHandler(svc service.ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

func (p *ProductHandler) CreateProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductCreateReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			logger.Info(err.Error())
			ctx.JSON(http.StatusOK, types.ErrorResponse(ctx, err))
			return
		}
		err := p.svc.CreateProduct(ctx, &req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, types.ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusCreated, types.RespSuccess(ctx, req))
	}
}
func (p *ProductHandler) UpdateProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductUpdateReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			logger.Info(err.Error())
			ctx.JSON(http.StatusOK, types.ErrorResponse(ctx, err))
			return
		}
		err := p.svc.UpdateProduct(ctx, &req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, types.ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccess(ctx, req))
	}
}
func (p *ProductHandler) DeleteProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := ctx.Param("uuid")
		err := p.svc.DeleteProduct(ctx, uuid)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, types.ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccess(ctx, nil))
	}
}
