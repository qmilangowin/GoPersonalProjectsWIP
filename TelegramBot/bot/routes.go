package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *BotApplication) routes() http.Handler {

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", app.home)
	// return secureHeaders(mux)

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))

	return secureHeaders(mux)
}
