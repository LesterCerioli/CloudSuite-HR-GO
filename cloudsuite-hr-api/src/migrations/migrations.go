package migrations

import (
	"cloudsuite-hr-api/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {

	err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
	if err != nil {
		return err
	}

	return db.AutoMigrate(&models.Time{})
}
