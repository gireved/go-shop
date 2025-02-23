package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type Product struct {
	Id               int64           `gorm:"primary_key;AUTO_INCREMENT"`
	Uuid             string          `gorm:"type:varchar(128);unique_index"`
	Name             string          `gorm:"type:varchar(128);unique_index"`
	Description      string          `gorm:"type:text"`
	Number           int             `gorm:"type:int"`
	OriginalPrice    decimal.Decimal `gorm:"type:decimal(12,2);"`
	PromotionalPrice decimal.Decimal `gorm:"type:decimal(12,2);"`
	Owner            string          `gorm:"type:varchar(128)"`
	Category         string          `gorm:"type:varchar(128)"`
	Status           bool            `gorm:"type:tinyint(1)"`
	CreatedAt        time.Time       `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;"`
	UpdatedAt        time.Time       `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
}
