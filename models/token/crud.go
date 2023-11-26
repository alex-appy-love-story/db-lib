package token

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func CreateToken(db *gorm.DB, cost float32) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&Token{
			Cost: decimal.NewFromFloat32(cost),
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
}

func GetToken(db *gorm.DB, ID uint) (*Token, error) {
	tok := &Token{}
	err := db.First(tok, ID).Error
	return tok, err
}
