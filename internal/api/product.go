package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-shop/internal/service"
	"go-shop/internal/types"
	"go-shop/pkg/logger"
	"strconv"

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
			ctx.JSON(http.StatusInternalServerError, types.ErrorResponse(ctx, 500, "参数校验失败", err))
			return
		}
		err := p.svc.CreateProduct(ctx, &req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, types.ErrorResponse(ctx, 500, "", err))
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
			ctx.JSON(http.StatusInternalServerError, types.ErrorResponse(ctx, 500, "参数校验未通过", err))
			return
		}
		err := p.svc.UpdateProduct(ctx, &req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, types.ErrorResponse(ctx, 500, "", err))
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
			ctx.JSON(http.StatusInternalServerError, types.ErrorResponse(ctx, 500, "", err))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccess(ctx, nil))
	}
}
func (p *ProductHandler) ListAllProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, _ := strconv.Atoi(ctx.Query("page"))
		pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
		res, err := p.svc.ListAllProducts(ctx, page, pageSize)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, types.ErrorResponse(ctx, 500, "", err))
			return
		}
		if res == nil || len(res) == 0 {
			err := errors.New("数据不存在")
			ctx.JSON(http.StatusNotFound, types.ErrorResponse(ctx, 404, "", err))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccess(ctx, res))
	}
}
func (p *ProductHandler) GetProductByNameHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Query("name")
		page, _ := strconv.Atoi(ctx.Query("page"))
		pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
		res, err := p.svc.GetProductByName(ctx, name, page, pageSize)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, types.ErrorResponse(ctx, 500, "", err))
			return
		}
		if res == nil || len(*res) == 0 {
			err := errors.New("数据不存在")
			ctx.JSON(http.StatusNotFound, types.ErrorResponse(ctx, 404, "", err))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccess(ctx, res))
	}
}
