package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	// _ "github.com/gin-gonic/gin"
	// 	create my sql connection
	// download go get github.com/lib/pq
	// then run go mod tidy and vendor
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "ppt"
	dbname   = "lms"
)

// const (
// 	// TODO fill this in directly or through environment variable
// 	// Build a DSN e.g. "postgres://username:password@url.com:5432/dbName"
// 	// "postgres://postgres@localhost:5436/test?sslmode=disable"
// 	db_dsn = "postgres://postgres:@localhost:5432/cslab?sslmode=disable"

// 	// postgres://postgres:abhinav@localhost:5432/cslab?sslmode=disable
// )

func connection_with_db() {
	db_dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", db_dsn) //postgres is the driver and db_dsn is the connector
	if err != nil {
		log.Fatal("Failed to open the DB connection", err)
	} else {
		fmt.Println("connected")
	}

}

// perfectly completed
