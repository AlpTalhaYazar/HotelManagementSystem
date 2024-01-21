package main

import (
	"github.com/alptalhayazar/hotelmanagementsystem/hotelmanagementservice/src/db"
	"github.com/alptalhayazar/hotelmanagementsystem/hotelmanagementservice/src/routes"
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

	router := routes.SetupRoutes(database)

	routerErr := router.Run(":8080")

	if routerErr != nil {
		log.Panicf("Error starting router: %v", routerErr)
	}
}
