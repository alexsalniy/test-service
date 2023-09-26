package sqlstore_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	godotenv.Load()
	databaseURL = os.Getenv("TESTDBURL")
	
	os.Exit(m.Run())
}