package order

import "gorm.io/gorm"

func CreateOrder(db *gorm.DB, order *OrderInfo) (*Order, error) {
	newOrder := &Order{
		OrderInfo: *order,
	}
	return newOrder, db.Create(newOrder).Error
}
