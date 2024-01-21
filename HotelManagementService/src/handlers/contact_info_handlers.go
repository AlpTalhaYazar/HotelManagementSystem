package handlers

import (
	"github.com/alptalhayazar/hotelmanagementsystem/hotelmanagementservice/src/models"
	"github.com/alptalhayazar/hotelmanagementsystem/hotelmanagementservice/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetContactInfos(context *gin.Context, db *gorm.DB) {
	var contactInfos []models.ContactInfo

	result := db.Find(&contactInfos)

	if result.Error != nil {
		response := utils.NewErrorResponse("Error retrieving contact infos!", utils.Error{Code: 500, Message: result.Error.Error()})

		context.JSON(http.StatusNotFound, response)
		return
	}

	response := utils.NewSuccessResponse("Successfully retrieved contact infos!", contactInfos)
	context.JSON(http.StatusOK, response)
}

func GetContactInfo(context *gin.Context, db *gorm.DB) {
	var contactInfo models.ContactInfo

	_, uuidParseError := uuid.Parse(context.Query("id"))

	if uuidParseError != nil {
		response := utils.NewErrorResponse("Error retrieving contact info!", utils.Error{Code: 400, Message: "Error parsing UUID!" + uuidParseError.Error()})

		context.JSON(http.StatusBadRequest, response)
		return
	}

	result := db.Where("id = ?", context.Query("id")).First(&contactInfo)

	if result.Error != nil {
		response := utils.NewErrorResponse("Error retrieving contact info!", utils.Error{Code: 404, Message: result.Error.Error()})

		context.JSON(http.StatusNotFound, response)
		return
	}

	response := utils.NewSuccessResponse("Successfully retrieved contact info!", contactInfo)
	context.JSON(http.StatusOK, response)
}

func CreateContactInfo(context *gin.Context, db *gorm.DB) {
	var contactInfo models.ContactInfo

	bindJSONError := context.BindJSON(&contactInfo)

	if bindJSONError != nil {
		response := utils.NewErrorResponse("Error creating contact info!", utils.Error{Code: 400, Message: bindJSONError.Error()})

		context.JSON(http.StatusBadRequest, response)
		return
	}

	dbCreateResult := db.Create(&contactInfo)

	if dbCreateResult.Error != nil {
		response := utils.NewErrorResponse("Error creating contact info!", utils.Error{Code: 400, Message: dbCreateResult.Error.Error()})

		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.NewSuccessResponse("Successfully created contact info!", contactInfo)
	context.JSON(http.StatusCreated, response)
}

func UpdateContactInfo(context *gin.Context, db *gorm.DB) {
	var contactInfo models.ContactInfo

	_, uuidParseError := uuid.Parse(context.Query("id"))

	if uuidParseError != nil {
		response := utils.NewErrorResponse("Error retrieving contact info!", utils.Error{Code: 400, Message: "Error parsing UUID!" + uuidParseError.Error()})

		context.JSON(http.StatusBadRequest, response)
		return
	}

	result := db.Where("id = ?", context.Query("id")).First(&contactInfo)

	if result.Error != nil {
		response := utils.NewErrorResponse("Error retrieving contact info!", utils.Error{Code: 404, Message: result.Error.Error()})

		context.JSON(http.StatusNotFound, response)
		return
	}

	bindJSONError := context.BindJSON(&contactInfo)

	if bindJSONError != nil {
		response := utils.NewErrorResponse("Error updating contact info!", utils.Error{Code: 400, Message: bindJSONError.Error()})

		context.JSON(http.StatusBadRequest, response)
		return
	}

	dbSaveResult := db.Save(&contactInfo)

	if dbSaveResult.Error != nil {
		response := utils.NewErrorResponse("Error updating contact info!", utils.Error{Code: 400, Message: dbSaveResult.Error.Error()})

		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.NewSuccessResponse("Successfully updated contact info!", contactInfo)
	context.JSON(http.StatusOK, response)
}

func DeleteContactInfo(context *gin.Context, db *gorm.DB) {
	var contactInfo models.ContactInfo

	_, uuidParseError := uuid.Parse(context.Query("id"))

	if uuidParseError != nil {
		response := utils.NewErrorResponse("Error deleting contact info!", utils.Error{Code: 400, Message: "Error parsing UUID!" + uuidParseError.Error()})

		context.JSON(http.StatusBadRequest, response)
		return
	}

	result := db.Where("id = ?", context.Query("id")).First(&contactInfo)

	if result.Error != nil {
		response := utils.NewErrorResponse("Error deleting contact info!", utils.Error{Code: 404, Message: "Record not found!"})

		context.JSON(http.StatusNotFound, response)
		return
	}

	dbDeleteResult := db.Delete(&contactInfo)

	if dbDeleteResult.Error != nil {
		response := utils.NewErrorResponse("Error deleting contact info!", utils.Error{Code: 500, Message: dbDeleteResult.Error.Error()})

		context.JSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.NewSuccessResponse("Successfully deleted contact info!", contactInfo)
	context.JSON(http.StatusOK, response)
}

func GetContactInfosByHotelId(context *gin.Context, db *gorm.DB) {
	var contactInfos []models.ContactInfo

	_, uuidParseError := uuid.Parse(context.Query("hotel_id"))

	if uuidParseError != nil {
		response := utils.NewErrorResponse("Error retrieving contact infos!", utils.Error{Code: 400, Message: "Error parsing UUID!" + uuidParseError.Error()})

		context.JSON(http.StatusBadRequest, response)
		return
	}

	result := db.Where("hotel_id = ?", context.Query("hotel_id")).Find(&contactInfos)

	if result.Error != nil {
		response := utils.NewErrorResponse("Error retrieving contact infos!", utils.Error{Code: 404, Message: result.Error.Error()})

		context.JSON(http.StatusNotFound, response)
		return
	}

	response := utils.NewSuccessResponse("Successfully retrieved contact infos!", contactInfos)
	context.JSON(http.StatusOK, response)
}
