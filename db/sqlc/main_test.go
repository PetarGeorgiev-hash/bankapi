package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/PetarGeorgiev-hash/bankapi/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Can't read env file from tests ", err.Error())
	}
	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	testQueries = New(testDb)
	os.Exit(m.Run())
}
