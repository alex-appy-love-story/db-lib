package user

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Creates a user with a starting balance of '1000.0'.
func CreateUser(db *gorm.DB, username string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&User{
			Username: username,
			Balance:  decimal.NewFromFloat32(1000.0),
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})
}

// Retrieve the user by username.
func GetUserByUsername(db *gorm.DB, username string) (*User, error) {
	usr := &User{}
	err := db.Where(&User{Username: username}).First(usr).Error
	return usr, err
}

// Retrieve the user by ID.
func GetUser(db *gorm.DB, ID uint) (*User, error) {
	usr := &User{}
	err := db.First(usr, ID).Error
	return usr, err
}
