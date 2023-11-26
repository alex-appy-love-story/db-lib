package inventory

import "gorm.io/gorm"

func CreateInventory(db *gorm.DB, info *InventoryInfo) (*Inventory, error) {
	inv := &Inventory{
		InventoryInfo: *info,
	}
	return inv, db.Create(inv).Error
}

func GetInventory(db *gorm.DB, tokenID uint) (*Inventory, error) {
	usr := &Inventory{}
	err := db.Where(&Inventory{
		InventoryInfo: InventoryInfo{
			TokenID: tokenID,
		},
	}).First(usr).Error
	return usr, err
}

func UpdateInventory(db *gorm.DB, info *InventoryInfo) (*Inventory, error) {
	var ret *Inventory = nil

	err := db.Transaction(func(tx *gorm.DB) error {

		var err error
		if ret, err = GetInventory(tx, info.TokenID); err != nil {
			return err
		}

		ret.Amount += info.Amount

		if err := tx.Save(&ret).Error; err != nil {
			return err
		}

		return nil
	})

	return ret, err
}

func AddInventory(db *gorm.DB, info *InventoryInfo) (*Inventory, error) {
	var ret *Inventory = nil

	err := db.Transaction(func(tx *gorm.DB) error {

		var exists bool = false
		err := tx.Model(&Inventory{}).
			Select("count(*) > 0").
			Where("token_id = ?", info.TokenID).
			Find(&exists).
			Error
		if err != nil {
			return err
		}

		if exists {
			// Entry exists, update it.
			if ret, err = UpdateInventory(tx, info); err != nil {
				return err
			}
		} else {
			// Entry does not exist, create it.
			if ret, err = CreateInventory(tx, info); err != nil {
				return err
			}
		}

		return nil
	})

	return ret, err
}
