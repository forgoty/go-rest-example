package store_test

import (
	"testing"
	"os"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "host=pg dbname=example_test sslmode=disable user=example password=example"
	}

	os.Exit(m.Run())
}
