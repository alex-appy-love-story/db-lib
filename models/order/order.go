package order

import "gorm.io/gorm"

type OrderInfo struct {
	UserID  uint `json:"user_id"`
	TokenID uint `json:"token_id"`
	Amount  uint `json:"amount"`
}

type Order struct {
	// ID
	// CreatedAt
	// UpdatedAt
	// DeletedAt
	gorm.Model

	// UserID
	// TokenID
	// Amount
	OrderInfo
}
