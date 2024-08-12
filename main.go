package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nmhoang2909/bank/api"
	db "github.com/nmhoang2909/bank/db/sqlc"
)

const (
	driverName    = "mysql"
	dataSource    = "root:secret@tcp/bank?parseTime=true"
	serverAddress = "localhost:8080"
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

	store := db.NewStore(conn)
	server := api.NewServer(store)
	if err := server.Start(serverAddress); err != nil {
		log.Fatal(err)
	}
}
