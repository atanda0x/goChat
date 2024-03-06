package main

import (
	"database/sql"
	"log"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "dbname=go-chat  sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
