package inventory

import "gorm.io/gorm"

type InventoryInfo struct {
	TokenID uint `json:"token_id"`
	Amount  uint `json:"amount"`
}

type Inventory struct {
	// ID
	// CreatedAt
	// UpdatedAt
	// DeletedAt
	gorm.Model

	// TokenID
	// Amount
	InventoryInfo
}
