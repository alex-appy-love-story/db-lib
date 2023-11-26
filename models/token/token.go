package token

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Token struct {
	// ID
	// CreatedAt
	// UpdatedAt
	// DeletedAt
	gorm.Model

	Cost decimal.Decimal `json:"cost" sql:"type:decimal(10,2);"`
}
