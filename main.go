package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	db "github.com/nmhoang2909/bank/db/sqlc"
)

const (
	driverName = "mysql"
	dataSource = "root:secret@tcp/bank?parseTime=true"
)

func main() {
	var err error
	conn, err := sql.Open(driverName, dataSource)
	if err != nil {
		log.Fatal(err)
	}

	if err = conn.Ping(); err != nil {
		log.Fatal(err)
	}

	db.NewStore(conn)

	os.Exit(m.Run())
}
