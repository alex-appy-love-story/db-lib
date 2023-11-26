package order

import "gorm.io/gorm"

type Order struct {
	// ID, CreatedAt, UpdatedAt, DeletedAt.
	gorm.Model

	// UserID.
	UserID uint `json:"user_id"`

	// Token amount.
	Amount uint `json:"amount"`
}
