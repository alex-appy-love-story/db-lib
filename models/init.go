// Migrate all tables.

package init

import (
	"log"

	"github.com/alex-appy-love-story/db-lib/models/inventory"
	"github.com/alex-appy-love-story/db-lib/models/order"
	"github.com/alex-appy-love-story/db-lib/models/token"
	"github.com/alex-appy-love-story/db-lib/models/user"
	"gorm.io/gorm"
)

func InitTables(db *gorm.DB) {
	if db == nil {
		log.Panic("Database is invalid.")
		return
	}

	migrator := db.Migrator()

	if !migrator.HasTable(&user.User{}) {
		err := db.AutoMigrate(&user.User{})
		if err != nil {
			log.Panic("Failed to migrate User table", err)
		}
	}

	if !migrator.HasTable(&token.Token{}) {
		err := db.AutoMigrate(&token.Token{})
		if err != nil {
			log.Panic("Failed to migrate Token table", err)
		}
	}

	if !migrator.HasTable(&order.Order{}) {
		err := db.AutoMigrate(&order.Order{})
		if err != nil {
			log.Panic("Failed to migrate Order table", err)
		}
	}

	if !migrator.HasTable(&inventory.Inventory{}) {
		err := db.AutoMigrate(&inventory.Inventory{})
		if err != nil {
			log.Panic("Failed to migrate Inventory table", err)
		}
	}
}
