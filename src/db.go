package main

import (
	"database/sql"
	"log"
	//_ "github.com/lib/pq"
)

var db *sql.DB

func InitializeDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	// Ensure the database connection is valid
	if err = db.Ping(); err != nil {
		return err
	}

	log.Println("Database connected")
	return nil
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}
