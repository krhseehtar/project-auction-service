package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection() (*sql.DB, error) {
	dbHost := "localhost"
	dbUser := "root"
	dbPassword := "password"
	dbName := "test"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Error pinging the database:", err)
		return nil, err
	}

	log.Println("Connected to the database successfully!")
	return db, nil
}
