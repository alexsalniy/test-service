package sqlstore

import (
	"database/sql"
	"strings"
	"testing"
)

func TestingDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if err := db.Ping(); err != nil {
			db.Exec("TRUNCATE %s CASCADE", strings.Join(tables, ", "))
		}

		db.Close()
	}
}