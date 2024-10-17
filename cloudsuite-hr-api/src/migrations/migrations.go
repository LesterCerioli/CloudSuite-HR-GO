package migrations

import (
	"cloudsuite-hr-api/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	// Cria a extensão uuid-ossp
	err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
	if err != nil {
		return err
	}

	// Cria a tabela times
	return db.AutoMigrate(&models.Time{})
}
