package dao

import (
	"context"
	"go-shop/internal/models"
	"go-shop/pkg/logger"
	"gorm.io/gorm"
)

type ProductDao interface {
	CreateProduct(ctx context.Context, product *models.Product) error
	GetProductsByName(ctx context.Context, name string, page, pageSize int) (*[]models.Product, error)
	GetAllProducts(ctx context.Context, page, pageSize int) ([]models.Product, error)
	UpdateProduct(ctx context.Context, updateProduct map[string]interface{}, uuid string) error
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

func (p *productDao) GetProductsByName(ctx context.Context, name string, page, pageSize int) (*[]models.Product, error) {
	var products []models.Product
	/*
		这个 result 包含了查询的执行状态，包括：
		查询是否成功。
		查询到的数据（会填充到传入的 &products 中）。
		如果有错误，则记录在 result.Error 中。
	*/
	result := p.DB.Model(models.Product{}).Where("name LIKE ?", name+"%").Scopes(paginate(page, pageSize)).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		logger.Info("Product not found")
	}
	return &products, nil
}

func (p *productDao) GetAllProducts(ctx context.Context, page, pageSize int) ([]models.Product, error) {
	var products []models.Product
	result := p.DB.Scopes(paginate(page, pageSize)).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		logger.Info("Product not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return products, result.Error
}

func (p *productDao) UpdateProduct(ctx context.Context, updateProduct map[string]interface{}, uuid string) error {
	//return p.DB.Model(&models.Product{}).Where("uuid = ?", product.Uuid).Updates(product).Error
	return p.DB.Model(&models.Product{}).Where("uuid = ?", uuid).Updates(updateProduct).Error
}

func (p *productDao) DeleteProduct(ctx context.Context, Uuid string) error {
	return p.DB.Where("uuid = ?", Uuid).Delete(&models.Product{}).Error
}
func paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page < 1 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Limit(pageSize).Offset(offset)
	}
}
