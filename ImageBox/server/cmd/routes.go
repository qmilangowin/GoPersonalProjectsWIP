package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *ServerApplication) routes() http.Handler {

	//use alice to chain the middleware
	//standardMiddleWare := alice.New(app.logRequest)

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/images", http.HandlerFunc(app.getGalleries))
	mux.Post("/images/upload", http.HandlerFunc(app.uploadImage))
	mux.Get("/images/single", http.HandlerFunc(app.getSingleImage))
	mux.Get("/images/faves", http.HandlerFunc(app.getFaves))
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", app.home)
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// return secureHeaders(mux)
	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
	//return standardMiddleWare.Then(mux)
}
