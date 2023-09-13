package initiliazer

import "gorm.io/gorm"

func SyncDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(user)
	
}