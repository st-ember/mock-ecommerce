package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB(connStr string) {
	var err error
	DB, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("Failed to connect to Database: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database unreachable: ", err)
	}
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
