package dao

import (
	"context"
	"go-shop/internal/models"
	"go-shop/pkg/logger"
	"gorm.io/gorm"
)

type ProductDao interface {
	CreateProduct(ctx context.Context, product *models.Product) error
	GetProductsByName(ctx context.Context, name string) (*[]models.Product, error)
	GetAllProducts(ctx context.Context) (*[]models.Product, error)
	UpdateProduct(ctx context.Context, product *models.Product) error
	DeleteProduct(ctx context.Context, Uuid string) error
}

type productDao struct {
	*gorm.DB
}

func NewProductDao(db *gorm.DB) ProductDao {
	return &productDao{DB: db}
}

func (p *productDao) CreateProduct(ctx context.Context, product *models.Product) error {
	return p.DB.Create(product).Error
}

func (p *productDao) GetProductsByName(ctx context.Context, name string) (*[]models.Product, error) {
	var products []models.Product
	/*
		这个 result 包含了查询的执行状态，包括：
		查询是否成功。
		查询到的数据（会填充到传入的 &products 中）。
		如果有错误，则记录在 result.Error 中。
	*/
	result := p.DB.Where("name LIKE ?", name+"%").Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		logger.Info("Product not found")
	}
	return &products, nil
}

func (p *productDao) GetAllProducts(ctx context.Context) (*[]models.Product, error) {
	var products []models.Product
	result := p.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		logger.Info("Product not found")
	}
	return &products, result.Error
}

func (p *productDao) UpdateProduct(ctx context.Context, product *models.Product) error {
	return p.DB.Save(product).Error
}

func (p *productDao) DeleteProduct(ctx context.Context, Uuid string) error {
	return p.DB.Delete(&models.Product{}, Uuid).Error
}
