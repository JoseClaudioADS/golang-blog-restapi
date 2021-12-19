package infra

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	Database *sqlx.DB
}

func (db *DB) Open() {
	databaseConnection, err := sqlx.Connect("postgres", "user=postgres password=postgres dbname=blog sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db.Database = databaseConnection
}

func (db *DB) Close() error {
	return db.Database.Close()
}
