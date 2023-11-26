package user

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Creates a user with a starting balance of '1000.0'.
func CreateUser(db *gorm.DB, username string) (*User, error) {
	usr := &User{
		UserInfo: UserInfo{
			Username: username,
			Balance:  decimal.NewFromFloat32(1000.0),
		},
	}

	return usr, db.Create(usr).Error
}

// Retrieve the user by username.
func GetUserByUsername(db *gorm.DB, username string) (*User, error) {
	usr := &User{}
	err := db.Where(&User{
		UserInfo: UserInfo{
			Username: username,
		},
	}).First(usr).Error
	return usr, err
}

// Retrieve the user by ID.
func GetUser(db *gorm.DB, ID uint) (*User, error) {
	usr := &User{}
	err := db.First(usr, ID).Error
	return usr, err
}
