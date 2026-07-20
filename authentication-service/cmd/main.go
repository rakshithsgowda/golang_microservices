package main

import (
	"authentication/data"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"
)

type Application struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("starting authentication service")

	// connecting to database

	app := Application{}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        app.routes(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}

func connetToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	for range 10 {
		db, err := sql.Open("pgx", dsn)
		if err != nil && db.Ping() == nil {
			log.Println("couldnot connect ot db")
			return db
		}
		log.Panicln("waiting for postgres")
		time.Sleep(time.Second * 2)
	}
	return nil
}
