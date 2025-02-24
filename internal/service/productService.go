package service

import (
	"context"
	"go-shop/internal/model"
	"go-shop/internal/repository/dao"
	"go-shop/pkg/logger"
)

type ProductService interface {
	GetProductByName(ctx context.Context, name string) (*[]model.Product, error)
	CreateProduct(ctx context.Context, product *model.Product) error
	UpdateProduct(ctx context.Context, product *model.Product) error
	DeleteProduct(ctx context.Context, uuid string) error
	ListAllProducts(ctx context.Context) (*[]model.Product, error)
}
type productService struct {
	productDao dao.ProductDao
}

func NewProductService(productDao dao.ProductDao) ProductService {
	return &productService{productDao: productDao}
}
func (s *productService) GetProductByName(ctx context.Context, name string) (*[]model.Product, error) {
	products, err := s.productDao.GetProductsByName(ctx, name)
	if err != nil {
		logger.Error("根据名称获取商品失败" + err.Error())
	}
	return products, err
}
func (s *productService) CreateProduct(ctx context.Context, product *model.Product) error {

	err := s.productDao.CreateProduct(ctx, product)
	if err != nil {
		logger.Error("新增商品出错" + err.Error())
		return err
	}
	return err
}
func (s *productService) UpdateProduct(ctx context.Context, product *model.Product) error {
	err := s.productDao.UpdateProduct(ctx, product)
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
func (s *productService) ListAllProducts(ctx context.Context) (*[]model.Product, error) {
	products, err := s.productDao.GetAllProducts(ctx)
	if err != nil {
		logger.Error("展示所有商品出错" + err.Error())
		return nil, err
	}
	return products, err
}
