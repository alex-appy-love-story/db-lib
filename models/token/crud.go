package token

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func CreateToken(db *gorm.DB, cost float32) (*Token, error) {
	tok := &Token{
		Cost: decimal.NewFromFloat32(cost),
	}
	return tok, db.Create(tok).Error
}

func GetToken(db *gorm.DB, ID uint) (*Token, error) {
	tok := &Token{}
	err := db.First(tok, ID).Error
	return tok, err
}
