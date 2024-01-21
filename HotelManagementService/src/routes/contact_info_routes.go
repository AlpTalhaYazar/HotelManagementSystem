package routes

import (
	"github.com/alptalhayazar/hotelmanagementsystem/hotelmanagementservice/src/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ContactInfoRoutes(router *gin.RouterGroup, db *gorm.DB) {
	contactInfoGroup := router.Group("/contact_infos")
	{
		contactInfoGroup.GET("", func(context *gin.Context) {
			if context.Query("id") != "" {
				handlers.GetContactInfo(context, db)
			} else if context.Query("hotel_id") != "" {
				handlers.GetContactInfosByHotelId(context, db)
			} else {
				handlers.GetContactInfos(context, db)
			}
		})
		contactInfoGroup.POST("", func(context *gin.Context) {
			handlers.CreateContactInfo(context, db)
		})
		contactInfoGroup.PUT("", func(context *gin.Context) {
			handlers.UpdateContactInfo(context, db)
		})
		contactInfoGroup.DELETE("", func(context *gin.Context) {
			handlers.DeleteContactInfo(context, db)
		})
	}
}
