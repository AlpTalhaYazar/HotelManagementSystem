package utils

import (
	"github.com/DATA-DOG/go-sqlmock"
	"net/http/httptest"
	"testing"
)

func RunTest(t *testing.T, ts *httptest.Server, route string, expectedStatusCode int, mock sqlmock.Sqlmock) {
	res, err := ts.Client().Get(ts.URL + route)

	if err != nil {
		t.Fatalf("Error sending request to server: %v", err)
	}

	if res.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, got %v", expectedStatusCode, res.StatusCode)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}
