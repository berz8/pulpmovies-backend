package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/libsql/libsql-client-go/libsql"
)

var DB *sql.DB

func Connect() {
	dbUrl := os.Getenv("DB")
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		log.Fatal("failed to connect to db", err)
	}
	DB = db
}
