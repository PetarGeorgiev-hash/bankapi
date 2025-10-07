package main

import (
	"database/sql"
	"log"

	"github.com/PetarGeorgiev-hash/bankapi/api"
	db "github.com/PetarGeorgiev-hash/bankapi/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	var err error
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Can't start server", err.Error())
	}
}
