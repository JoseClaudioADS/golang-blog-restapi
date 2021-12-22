package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JoseClaudioADS/golang-blog-restapi/app"
	"github.com/JoseClaudioADS/golang-blog-restapi/app/infra/database"
	"github.com/joho/godotenv"
)

const SERVER_PORT = "8080"

func main() {

	loadEnv()

	db := database.DB{}
	db.Open()
	defer db.Close()
	db.Connection.MustExec(database.DatabaseSchema)

	server := app.New(&db)

	http.Handle("/", server.Router)

	err := http.ListenAndServe(fmt.Sprintf(":%s", SERVER_PORT), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
