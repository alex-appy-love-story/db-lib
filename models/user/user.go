package user

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type UserInfo struct {
	Username string          `json:"username"`
	Balance  decimal.Decimal `json:"balance" sql:"type:decimal(10,2);"`
}

type User struct {
	// ID
	// CreatedAt
	// UpdatedAt
	// DeletedAt
	gorm.Model

	// Username
	// Balance
	UserInfo
}
