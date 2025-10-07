package main

import (
	"database/sql"
	"log"

	"github.com/PetarGeorgiev-hash/bankapi/api"
	db "github.com/PetarGeorgiev-hash/bankapi/db/sqlc"
	"github.com/PetarGeorgiev-hash/bankapi/util"
	_ "github.com/lib/pq"
)

func main() {
	var err error
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't read env file", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can't start server", err.Error())
	}
}
