package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JoseClaudioADS/golang-blog-restapi/app"
)

const SERVER_PORT = "8080"

func main() {
	app := app.New()

	http.Handle("/", app.Router)

	err := http.ListenAndServe(fmt.Sprintf(":%s", SERVER_PORT), nil)
	if err != nil {
		log.Fatal(err)
	}
}
