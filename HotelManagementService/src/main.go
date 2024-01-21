package main

import (
	"github.com/alptalhayazar/hotelmanagementsystem/hotelmanagementservice/src/db"
	"log"
)

func main() {
	database, databaseErr := db.ConnectDatabase("Default")

	if databaseErr != nil {
		log.Panicf("Error connecting to database: %v", databaseErr)
	}

	databaseMigrationError := db.MigrateDatabase(database)
	if databaseMigrationError != nil {
		log.Panicf("Error migrating database: %v", databaseMigrationError)
	}
}
