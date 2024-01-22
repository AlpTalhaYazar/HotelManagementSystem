package handlers

import (
	"fmt"
	"github.com/alptalhayazar/hotelmanagementsystem/hotelmanagementservice/src/models"
	"github.com/alptalhayazar/hotelmanagementsystem/hotelmanagementservice/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetHotels(context *gin.Context, db *gorm.DB) {
	var hotels []models.Hotel

	result := db.Find(&hotels)

	if result.Error != nil {
		response := utils.NewErrorResponse("Error retrieving hotels!", utils.Error{Code: 500, Message: result.Error.Error()})

		context.JSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.NewSuccessResponse("Successfully retrieved hotels!", hotels)
	context.JSON(http.StatusOK, response)
}

func GetHotel(context *gin.Context, db *gorm.DB) {
	var hotel models.Hotel

	_, uuidParseError := uuid.Parse(context.Query("id"))

	if uuidParseError != nil {
		response := utils.NewErrorResponse("Error retrieving hotel!", utils.Error{Code: 400, Message: fmt.Sprintf("Error parsing UUID: %s", uuidParseError.Error())})

		context.JSON(http.StatusBadRequest, response)
		return
	}

	result := db.Where("\"Id\" = ?", context.Query("id")).First(&hotel)

	if result.Error != nil {
		response := utils.NewErrorResponse("Error retrieving hotel!", utils.Error{Code: 404, Message: result.Error.Error()})

		context.JSON(http.StatusNotFound, response)
		return
	}

	response := utils.NewSuccessResponse("Successfully retrieved hotel!", hotel)
	context.JSON(http.StatusOK, response)
}

func CreateHotel(context *gin.Context, db *gorm.DB) {
	var hotel models.Hotel

	err := context.BindJSON(&hotel)

	if err != nil {
		response := utils.NewErrorResponse("Error creating hotel!", utils.Error{Code: 400, Message: err.Error()})

		context.JSON(http.StatusBadRequest, response)
		return
	}

	hotel.Id = uuid.New()

	result := db.Create(&hotel)

	if result.Error != nil {
		response := utils.NewErrorResponse("Error creating hotel!", utils.Error{Code: 500, Message: result.Error.Error()})

		context.JSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.NewSuccessResponse("Successfully created hotel!", hotel)
	context.JSON(http.StatusCreated, response)
}

func UpdateHotel(context *gin.Context, db *gorm.DB) {
	var hotel models.Hotel

	fmt.Println("put")

	_, uuidParseError := uuid.Parse(context.Query("id"))

	if uuidParseError != nil {
		response := utils.NewErrorResponse("Error retrieving hotel!", utils.Error{Code: 400, Message: fmt.Sprintf("Error parsing UUID: %s", uuidParseError.Error())})

		context.JSON(http.StatusBadRequest, response)
		return
	}

	result := db.Where("Id = ?", context.Query("id")).First(&hotel)

	if result.Error != nil {
		response := utils.NewErrorResponse("Error retrieving hotel!", utils.Error{Code: 404, Message: result.Error.Error()})

		context.JSON(http.StatusNotFound, response)
		return
	}

	err := context.BindJSON(&hotel)

	if err != nil {
		response := utils.NewErrorResponse("Error updating hotel!", utils.Error{Code: 400, Message: err.Error()})

		context.JSON(http.StatusBadRequest, response)
		return
	}

	dbSaveResult := db.Save(&hotel)

	if dbSaveResult.Error != nil {
		response := utils.NewErrorResponse("Error updating hotel!", utils.Error{Code: 500, Message: dbSaveResult.Error.Error()})

		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.NewSuccessResponse("Successfully updated hotel!", hotel)
	context.JSON(http.StatusOK, response)
}

func DeleteHotel(context *gin.Context, db *gorm.DB) {
	var hotel models.Hotel

	_, uuidParseError := uuid.Parse(context.Query("id"))

	if uuidParseError != nil {
		response := utils.NewErrorResponse("Error retrieving hotel!", utils.Error{Code: 400, Message: fmt.Sprintf("Error parsing UUID: %s", uuidParseError.Error())})

		context.JSON(http.StatusBadRequest, response)
		return
	}

	result := db.Where("Id = ?", context.Query("id")).First(&hotel)

	if result.Error != nil {
		response := utils.NewErrorResponse("Error retrieving hotel!", utils.Error{Code: 404, Message: result.Error.Error()})

		context.JSON(http.StatusNotFound, response)
		return
	}

	dbDeleteResult := db.Delete(&hotel)

	if dbDeleteResult.Error != nil {
		response := utils.NewErrorResponse("Error deleting hotel!", utils.Error{Code: 500, Message: dbDeleteResult.Error.Error()})

		context.JSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.NewSuccessResponse("Successfully deleted hotel!", hotel)
	context.JSON(http.StatusOK, response)
}
