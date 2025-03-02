package service

import (
	"context"
	"github.com/google/uuid"
	"go-shop/internal/models"
	"go-shop/internal/repository/dao"
	"go-shop/internal/types"
	"go-shop/pkg/logger"
	"time"
)

type ProductService interface {
	GetProductByName(ctx context.Context, name string) (*[]models.Product, error)
	CreateProduct(ctx context.Context, product *types.ProductCreateReq) error
	UpdateProduct(ctx context.Context, product *types.ProductUpdateReq) error
	DeleteProduct(ctx context.Context, uuid string) error
	ListAllProducts(ctx context.Context) (*[]models.Product, error)
}
type productService struct {
	productDao dao.ProductDao
}

func NewProductService(productDao dao.ProductDao) ProductService {
	return &productService{productDao: productDao}
}
func (s *productService) GetProductByName(ctx context.Context, name string) (*[]models.Product, error) {
	products, err := s.productDao.GetProductsByName(ctx, name)
	if err != nil {
		logger.Error("根据名称获取商品失败" + err.Error())
	}
	return products, err
}
func (s *productService) CreateProduct(ctx context.Context, req *types.ProductCreateReq) error {

	err := s.productDao.CreateProduct(ctx, reqToProduct(req))
	if err != nil {
		logger.Error("新增商品出错" + err.Error())
		return err
	}
	return err
}
func (s *productService) UpdateProduct(ctx context.Context, req *types.ProductUpdateReq) error {
	updates := make(map[string]interface{})
	if req.Name != nil {
		updates["name"] = *req.Name
	}

	if req.Number != nil {
		updates["number"] = *req.Number
	}

	if req.Status != nil {
		updates["status"] = *req.Status
	}

	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.Owner != nil {
		updates["owner"] = *req.Owner
	}
	if req.Category != nil {
		updates["category"] = *req.Category
	}
	if req.PromotionPrice != nil {
		updates["promotionPrice"] = *req.PromotionPrice
	}
	if req.OriginalPrice != nil {
		updates["originalPrice"] = *req.OriginalPrice
	}
	err := s.productDao.UpdateProduct(ctx, updates, req.Uuid)
	if err != nil {
		logger.Error("更新商品时出错" + err.Error())
		return err
	}
	return err
}
func (s *productService) DeleteProduct(ctx context.Context, uuid string) error {
	err := s.productDao.DeleteProduct(ctx, uuid)
	if err != nil {
		logger.Error("删除商品时出错" + err.Error())
		return err
	}
	return err

}
func (s *productService) ListAllProducts(ctx context.Context) (*[]models.Product, error) {
	products, err := s.productDao.GetAllProducts(ctx)
	if err != nil {
		logger.Error("展示所有商品出错" + err.Error())
		return nil, err
	}
	return products, err
}

func reqToProduct(req *types.ProductCreateReq) *models.Product {
	uid := uuid.New().String()
	return &models.Product{
		Name:             req.Name,
		Uuid:             uid,
		Description:      req.Description,
		Number:           req.Number,
		OriginalPrice:    req.OriginalPrice,
		PromotionalPrice: req.PromotionPrice,
		Owner:            req.Owner,
		Category:         req.Category,
		Status:           req.Status,
		CreatedAt:        time.Now(),
	}
}
