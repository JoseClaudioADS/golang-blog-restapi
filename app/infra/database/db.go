package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	Connection *sqlx.DB
}

func (db *DB) Open() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	connectionInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, database, host, port)

	databaseConnection, err := sqlx.Connect("postgres", connectionInfo)
	if err != nil {
		log.Fatalln(err)
	}
	db.Connection = databaseConnection
}

func (db *DB) Close() error {
	return db.Connection.Close()
}
