package types

import "github.com/shopspring/decimal"

type ProductCreateReq struct {
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Number         int             `json:"number"`
	OriginalPrice  decimal.Decimal `json:"price"`
	PromotionPrice decimal.Decimal `json:"promotion"`
	Owner          string          `json:"owner"`
	Category       string          `json:"category"`
	Status         bool            `json:"status"`
}
type ProductUpdateReq struct {
	Uuid           string           `json:"uuid"`
	Name           *string          `json:"name"`
	Description    *string          `json:"description"`
	Number         *int             `json:"number"`
	OriginalPrice  *decimal.Decimal `json:"price"`
	PromotionPrice *decimal.Decimal `json:"promotion"`
	Owner          *string          `json:"owner"`
	Category       *string          `json:"category"`
	Status         *bool            `json:"status"`
}
