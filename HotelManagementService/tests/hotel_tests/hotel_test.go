package hotel_tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alptalhayazar/hotelmanagementsystem/hotelmanagementservice/tests/Database"
	"github.com/alptalhayazar/hotelmanagementsystem/hotelmanagementservice/tests/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"testing"
)

func TestGetHotels(t *testing.T) {
	gin.SetMode(gin.TestMode)

	_, mock, ts, teardown, err := Database.SetupMockDatabase()
	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}
	defer teardown()

	rows := sqlmock.NewRows([]string{"Id", "Name", "Description", "City", "Country", "Stars"})
	rows.AddRow("1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6", "Hotel 1", "Description 1", "City 1", "Country 1", 5)
	rows.AddRow("2fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6", "Hotel 2", "Description 2", "City 2", "Country 2", 4)
	mock.ExpectQuery("^SELECT \\* FROM \"Hotels\"$").WillReturnRows(rows)

	utils.RunTest(t, ts, "/api/hotels", 200, mock)
}

func TestGetHotelById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	_, mock, ts, teardown, err := Database.SetupMockDatabase()
	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}
	defer teardown()

	rows := sqlmock.NewRows([]string{"Id", "Name", "Description", "City", "Country", "Stars"})
	rows.AddRow("1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6", "Hotel 1", "Description 1", "City 1", "Country 1", 5)
	mock.ExpectQuery("^SELECT \\* FROM \"Hotels\" WHERE \"Id\" = \\$1 ORDER BY \"Hotels\".\"Id\" LIMIT 1$").WithArgs("1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6").WillReturnRows(rows)

	utils.RunTest(t, ts, "/api/hotels?id=1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6", 200, mock)
}

func TestGetHotelByIdNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	_, mock, ts, teardown, err := Database.SetupMockDatabase()
	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}
	defer teardown()

	mock.ExpectQuery("^SELECT \\* FROM \"Hotels\" WHERE \"Id\" = \\$1 ORDER BY \"Hotels\".\"Id\" LIMIT 1$").WithArgs("1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6").WillReturnError(gorm.ErrRecordNotFound)

	utils.RunTest(t, ts, "/api/hotels?id=1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6", 404, mock)
}

func TestGetHotelByIdInvalidUUID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	_, mock, ts, teardown, err := Database.SetupMockDatabase()
	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}
	defer teardown()

	utils.RunTest(t, ts, "/api/hotels?id=invalid-uuid", 400, mock)
}
