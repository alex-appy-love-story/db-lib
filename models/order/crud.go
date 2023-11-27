package order

import (
	"fmt"

	"gorm.io/gorm"
)

func CreateOrder(db *gorm.DB, order *OrderInfo) (*Order, error) {
	newOrder := &Order{
		OrderInfo:   *order,
		OrderStatus: PENDING,
	}
	return newOrder, db.Create(newOrder).Error
}

func GetOrder(db *gorm.DB, orderID uint) (*Order, error) {
	ord := &Order{}
	err := db.First(ord, orderID).Error
	return ord, err
}

func SetOrderStatus(db *gorm.DB, orderID uint, orderStatus OrderStatus) (*Order, error) {
	var ord *Order = nil
	var err error = nil

	err = db.Transaction(func(tx *gorm.DB) error {

		if ord, err = GetOrder(tx, orderID); err != nil {
			return err
		}

		ord.OrderStatus = orderStatus

		if err = tx.Save(ord).Error; err != nil {
			return err
		}

		return nil
	})

	return ord, err
}

// Only set if a value isn't set yet.
func SetOrderStatusFail(db *gorm.DB, orderID uint, orderStatus OrderStatus) (*Order, error) {
	var ord *Order = nil
	var err error = nil

	err = db.Transaction(func(tx *gorm.DB) error {

		if ord, err = GetOrder(tx, orderID); err != nil {
			return err
		}

		if ord.OrderStatus != PENDING {
			return fmt.Errorf("aborted, value already set")
		}

		ord.OrderStatus = orderStatus

		if err = tx.Save(ord).Error; err != nil {
			return err
		}

		return nil
	})

	return ord, err
}

func DeleteOrder(db *gorm.DB, orderID uint) error {
	return db.Delete(&Order{}, orderID).Error
}
