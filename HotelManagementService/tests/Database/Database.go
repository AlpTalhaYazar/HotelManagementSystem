package Database

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alptalhayazar/hotelmanagementsystem/hotelmanagementservice/src/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http/httptest"
)

func SetupMockDatabase() (*gin.Engine, sqlmock.Sqlmock, *httptest.Server, func(), error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: database}), &gorm.Config{})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	router := gin.Default()
	router = routes.SetupRoutes(gormDB)
	ts := httptest.NewServer(router)

	teardown := func() {
		ts.Close()
		database.Close()
	}

	return router, mock, ts, teardown, nil
}
