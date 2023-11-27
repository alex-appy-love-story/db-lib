package user

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type User struct {
	// ID
	// CreatedAt
	// UpdatedAt
	// DeletedAt
	gorm.Model

	// Username
	Balance decimal.Decimal `json:"balance" sql:"type:decimal(10,2);"`
}
