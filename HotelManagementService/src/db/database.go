package db

import (
	"fmt"
	"github.com/alptalhayazar/hotelmanagementsystem/hotelmanagementservice/src/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"slices"
)

var DB *gorm.DB

func ConnectDatabase(databaseName string) (*gorm.DB, error) {
	configuration := utils.ReadConfiguration()
	index := slices.IndexFunc(configuration.Databases, func(connection utils.Database) bool { return connection.Name == databaseName })
	connectionString := configuration.Databases[index].ConnectionString

	database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatalf("Database connection failed. Database error: %s", err.Error())
	}

	DB = database

	fmt.Println("Successfully connected to database!")

	return DB, err
}
