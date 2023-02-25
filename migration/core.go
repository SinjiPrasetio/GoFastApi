package migration

import "gorm.io/gorm"

// Define your auto migration here.
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
