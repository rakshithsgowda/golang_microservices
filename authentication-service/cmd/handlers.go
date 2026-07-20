package main

import "net/http"

func (a *Application) Auth(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email string 
		Password string
	}

	w.a.
}
