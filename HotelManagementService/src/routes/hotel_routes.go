package routes

import (
	"github.com/alptalhayazar/hotelmanagementsystem/hotelmanagementservice/src/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HotelRoutes(router *gin.RouterGroup, db *gorm.DB) {
	hotelGroup := router.Group("/hotels")
	{
		hotelGroup.GET("", func(context *gin.Context) {
			if context.Query("id") != "" {
				handlers.GetHotel(context, db)
			} else {
				handlers.GetHotels(context, db)
			}
		})
		hotelGroup.POST("", func(context *gin.Context) {
			handlers.CreateHotel(context, db)
		})
		hotelGroup.PUT("", func(context *gin.Context) {
			handlers.UpdateHotel(context, db)
		})
		hotelGroup.DELETE("", func(context *gin.Context) {
			handlers.DeleteHotel(context, db)
		})
	}
}
