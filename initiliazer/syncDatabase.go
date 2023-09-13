package initiliazer

import (
	"inventory-api/model"

	"gorm.io/gorm"
)

func SyncDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(model.User{}, model.Supplier{}, model.Transaction{}, model.Product{})
	return err
}
