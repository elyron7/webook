package dao

import "gorm.io/gorm"

func InitTable(db *gorm.DB) error {
	// db.Migrator().DropTable(&User{})

	return db.AutoMigrate(&User{})
}
