package contact_info_tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alptalhayazar/hotelmanagementsystem/hotelmanagementservice/tests/Database"
	"github.com/alptalhayazar/hotelmanagementsystem/hotelmanagementservice/tests/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"testing"
)

func TestGetContactInfos(t *testing.T) {
	gin.SetMode(gin.TestMode)

	_, mock, ts, teardown, err := Database.SetupMockDatabase()

	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}

	defer teardown()

	rows := sqlmock.NewRows([]string{"Id", "HotelId", "Name", "Phone", "Email"})
	rows.AddRow("1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6", "1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6", "Contact Info 1", "123456789", "qwe@qwe.com")
	rows.AddRow("2fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6", "2fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6", "Contact Info 2", "987654321", "asd@asd.com")
	mock.ExpectQuery("^SELECT \\* FROM \"ContactInfos\"$").WillReturnRows(rows)

	utils.RunTest(t, ts, "/api/contact_infos", 200, mock)
}

func TestGetContactInfoById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	_, mock, ts, teardown, err := Database.SetupMockDatabase()

	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}

	defer teardown()

	rows := sqlmock.NewRows([]string{"Id", "HotelId", "Name", "Phone", "Email"})
	rows.AddRow("1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6", "1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6", "Contact Info 1", "123456789", "qwe@qwe.com")
	mock.ExpectQuery("^SELECT \\* FROM \"ContactInfos\" WHERE Id = \\$1 ORDER BY \"ContactInfos\".\"Id\" LIMIT 1$").WithArgs("1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6").WillReturnRows(rows)

	utils.RunTest(t, ts, "/api/contact_infos?id=1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6", 200, mock)
}

func TestGetContactInfoByHotelId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	_, mock, ts, teardown, err := Database.SetupMockDatabase()

	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}

	defer teardown()

	rows := sqlmock.NewRows([]string{"Id", "HotelId", "Name", "Phone", "Email"})
	rows.AddRow("1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6", "1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6", "Contact Info 1", "123456789", "qwe@qwe.com")
	mock.ExpectQuery("^SELECT \\* FROM \"ContactInfos\" WHERE HotelId = \\$1$").WithArgs("1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6").WillReturnRows(rows)

	utils.RunTest(t, ts, "/api/contact_infos?hotel_id=1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6", 200, mock)
}

func TestGetContactInfoByIdNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	_, mock, ts, teardown, err := Database.SetupMockDatabase()

	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}

	defer teardown()

	mock.ExpectQuery("^SELECT \\* FROM \"ContactInfos\" WHERE Id = \\$1 ORDER BY \"ContactInfos\".\"Id\" LIMIT 1$").WithArgs("1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6").WillReturnError(gorm.ErrRecordNotFound)

	utils.RunTest(t, ts, "/api/contact_infos?id=1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6", 404, mock)
}

func TestGetContactInfoByHotelIdNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	_, mock, ts, teardown, err := Database.SetupMockDatabase()

	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}

	defer teardown()

	mock.ExpectQuery("^SELECT \\* FROM \"ContactInfos\" WHERE HotelId = \\$1$").WithArgs("1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6").WillReturnError(gorm.ErrRecordNotFound)

	utils.RunTest(t, ts, "/api/contact_infos?hotel_id=1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb6", 404, mock)
}

func TestGetContactInfoByIdInvalidUUID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	_, mock, ts, teardown, err := Database.SetupMockDatabase()

	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}

	defer teardown()

	utils.RunTest(t, ts, "/api/contact_infos?id=1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb", 400, mock)
}

func TestGetContactInfoByHotelIdInvalidUUID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	_, mock, ts, teardown, err := Database.SetupMockDatabase()

	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}

	defer teardown()

	utils.RunTest(t, ts, "/api/contact_infos?hotel_id=1fe3b3f0-bc7c-4587-b1ab-c8d4b8304fb", 400, mock)
}
