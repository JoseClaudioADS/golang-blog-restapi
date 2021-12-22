package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JoseClaudioADS/golang-blog-restapi/app"
	"github.com/JoseClaudioADS/golang-blog-restapi/app/infra/database"
)

const SERVER_PORT = "8080"

func main() {

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
