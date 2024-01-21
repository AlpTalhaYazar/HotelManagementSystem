package main

import (
	"github.com/alptalhayazar/hotelmanagementsystem/hotelmanagementservice/src/db"
	"log"
)

func main() {
	_, databaseErr := db.ConnectDatabase("Default")

	if databaseErr != nil {
		log.Panicf("Error connecting to database: %v", databaseErr)
	}
}
