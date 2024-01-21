package db

import (
	"github.com/alptalhayazar/hotelmanagementsystem/hotelmanagementservice/src/models"
	"gorm.io/gorm"
	"log"
)

func MigrateDatabase(db *gorm.DB) error {
	err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
	if err != nil {
		return err
	}

	hotelModelMigrationError := db.AutoMigrate(&models.Hotel{})
	if hotelModelMigrationError != nil {
		log.Fatalf("Error migrating hotel model: %v", hotelModelMigrationError)
		return hotelModelMigrationError
	}

	contactInformationModelMigrationError := db.AutoMigrate(&models.ContactInfo{})
	if contactInformationModelMigrationError != nil {
		log.Fatalf("Error migrating contact information model: %v", contactInformationModelMigrationError)
		return contactInformationModelMigrationError
	}

	return nil
}
