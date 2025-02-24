package dao

import "gorm.io/gorm"

type ProductDAO interface {
}

type ProductDao struct {
	*gorm.DB
}
