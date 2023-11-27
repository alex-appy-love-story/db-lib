package order

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type OrderStatus string

const (
	PENDING                   OrderStatus = "PENDING"
	SUCCESS                   OrderStatus = "SUCCESS"
	FAIL                      OrderStatus = "FAIL"
	PAYMENT_FAIL_INSUFFICIENT OrderStatus = "PAYMENT_FAIL_INSUFFICIENT"
	INVENTORY_FAIL_STOCK      OrderStatus = "INVENTORY_FAIL_STOCK"
)

func (self *OrderStatus) Scan(value interface{}) error {
	*self = OrderStatus(value.([]byte))
	return nil
}

func (self OrderStatus) Value() (driver.Value, error) {
	return string(self), nil
}

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

	OrderStatus OrderStatus `json:"order_status" sql:"type:ENUM('PENDING', 'SUCCESS', 'FAIL', 'PAYMENT_FAIL_INSUFFICIENT', 'INVENTORY_FAIL_STOCK')" gorm:"column:order_status"`
}
