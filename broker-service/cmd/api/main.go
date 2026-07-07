package main

import (
	"fmt"
	"log"
	"net/http"
)

const webport = "8080"

type Application struct{}

func main() {
	app := Application{}
	log.Printf("starting broker service...!!! on the port %s", webport)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webport),
		Handler: app.routes(),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
