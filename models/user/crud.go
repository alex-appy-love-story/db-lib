package user

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Creates a user with a starting balance of '1000.0'.
func CreateUser(db *gorm.DB, username string) (*User, error) {
	usr := &User{
		Username: username,
		Balance:  decimal.NewFromFloat32(1000.0),
	}

	return usr, db.Create(usr).Error
}

// Retrieve the user by ID.
func GetUser(db *gorm.DB, ID uint) (*User, error) {
	usr := &User{}
	err := db.First(usr, ID).Error
	return usr, err
}

// Retrieve the user by username.
func GetUserByUsername(db *gorm.DB, username string) (*User, error) {
	usr := &User{}
	err := db.Where(&User{
		Username: username,
	}).First(usr).Error
	return usr, err
}

func UpdateUserBalance(db *gorm.DB, ID uint, value decimal.Decimal) (*User, error) {
	var ret *User = nil

	err := db.Transaction(func(tx *gorm.DB) error {

		var err error
		if ret, err = GetUser(tx, ID); err != nil {
			return err
		}

		ret.Balance = value

		if err := tx.Save(&ret).Error; err != nil {
			return err
		}

		return nil
	})

	return ret, err
}
