package db_test

import (
	"database/sql"
	db "github/eskye/fingreat_backend/db/sqlc"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQuery *db.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open("postgres", "postgres://root:secret@localhost:5432/fingreat_db?sslmode=disable")
	if err != nil {
		log.Fatal("Could not connect to database", err)
	}
	testQuery = db.New(conn)

	os.Exit(m.Run())
}
