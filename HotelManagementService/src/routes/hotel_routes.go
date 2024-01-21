package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HotelRoutes(router *gin.RouterGroup, db *gorm.DB) {
	hotelGroup := router.Group("/hotels")
	{
		hotelGroup.GET("", func(context *gin.Context) {
			if context.Query("id") != "" {
				fmt.Printf("ID: %s\n", context.Query("id"))
			} else {
				fmt.Println("ID: not found")
			}
		})
		hotelGroup.POST("", func(context *gin.Context) {
			fmt.Println("POST")
		})
		hotelGroup.PUT("", func(context *gin.Context) {
			fmt.Println("PUT")
		})
		hotelGroup.DELETE("", func(context *gin.Context) {
			fmt.Println("DELETE")
		})
	}
}
