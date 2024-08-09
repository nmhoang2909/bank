package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const (
	driverName = "mysql"
	dataSource = "root:secret@tcp/bank?parseTime=true"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(driverName, dataSource)
	if err != nil {
		log.Fatal(err)
	}

	if err = testDB.Ping(); err != nil {
		log.Fatal(err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
