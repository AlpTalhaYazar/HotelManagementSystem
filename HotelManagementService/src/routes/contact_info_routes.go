package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ContactInfoRoutes(router *gin.RouterGroup, db *gorm.DB) {
	contactInfoGroup := router.Group("/contact_infos")
	{
		contactInfoGroup.GET("", func(context *gin.Context) {
			if context.Query("id") != "" {
				fmt.Printf("ID: %s\n", context.Query("id"))
			} else if context.Query("hotel_id") != "" {
				fmt.Printf("Hotel ID: %s\n", context.Query("hotel_id"))
			} else {
				fmt.Println("ID: not found")
			}
		})
		contactInfoGroup.POST("", func(context *gin.Context) {
			fmt.Println("POST")
		})
		contactInfoGroup.PUT("", func(context *gin.Context) {
			fmt.Println("PUT")
		})
		contactInfoGroup.DELETE("", func(context *gin.Context) {
			fmt.Println("DELETE")
		})
	}
}
