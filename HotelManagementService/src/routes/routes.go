package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		HotelRoutes(api, db)
		ContactInfoRoutes(api, db)
	}

	return router
}
