package pgdb

import (
	"database/sql"
	"fmt"

	"github.com/abhishek2966/singleton-demo/pkg/config"
	_ "github.com/lib/pq"
)

var instance *sql.DB

func GetDBInstance() (*sql.DB, error) {
	var err error
	if instance == nil {
		instance, err = openConnection()
		return instance, err
	}
	return instance, err
}

func openConnection() (*sql.DB, error) {
	fmt.Println("connecting")
	host := config.ReadConfig("pgdb", "host")
	port := config.ReadConfig("pgdb", "port")
	user := config.ReadConfig("pgdb", "user")
	password := config.ReadConfig("pgdb", "password")
	dbname := config.ReadConfig("pgdb", "dbname")
	dataSourceName := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dataSourceName)
	return db, err
}
