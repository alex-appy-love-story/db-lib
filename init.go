// Migrate all tables.

package init

import (
	"fmt"

	"gorm.io/gorm"
)

func InitTables(db *gorm.DB, table interface{}) error {
	if db == nil {
		return fmt.Errorf("Database is invalid.")
	}

	migrator := db.Migrator()

	if !migrator.HasTable(table) {
		err := db.AutoMigrate(table)
		if err != nil {
			return err
		}
	}

	return nil
}
