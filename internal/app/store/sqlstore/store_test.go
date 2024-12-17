package sqlstore_test

import (
	"os"
	"testing"
)

var (
	database_url string
)

func TestMain(m *testing.M) {
	database_url = os.Getenv("DATABASE_URL")
	if database_url == "" {
		database_url = "host=localhost dbname=restapi_test user=postgres password=pwd123 sslmode=disable"
	}

	os.Exit(m.Run())
}
