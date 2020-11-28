package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func GetDB() *sql.DB {
	var err error

	if db == nil {
		connStr := "user=webserviceuser password=webserviceuserpwd dbname=firstwebservice sslmode=disable"
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
	}
	return db
}
