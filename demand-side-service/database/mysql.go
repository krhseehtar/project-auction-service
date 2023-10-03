package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection() (*sql.DB, error) {
	dbHost := os.Getenv("MYSQL_HOST")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_NAME")
	dbPort := os.Getenv("MYSQL_PORT")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	log.Println("database connection string:", dataSourceName)

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
