package main

import (
	"log"
	"net/http"
)

func (a *Application) handleSubmission(w http.ResponseWriter, r *http.Request) {
	log.Println("hi")
}
