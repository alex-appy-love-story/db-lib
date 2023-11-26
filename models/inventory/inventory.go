package inventory

import "gorm.io/gorm"

type Inventory struct {
	// ID, CreatedAt, UpdatedAt, DeletedAt.
	gorm.Model

	// UserID.
	TokenID uint `json:"token_id"`

	// In stock amount.
	Amount uint
}
