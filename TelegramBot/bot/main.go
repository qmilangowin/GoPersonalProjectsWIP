package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type BotApplication struct{}

func main() {

	app := &BotApplication{}

	srv := &http.Server{
		Handler:      app.routes(),
		Addr:         ":8080",
		WriteTimeout: 20 * time.Second,
		ReadTimeout:  20 * time.Second,
	}

	//start the server
	go func() {
		fmt.Println("Starting server on port", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println("could not start server", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	log.Println("Received Terminate, graceful shutdown %s", sig)
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv.Shutdown(tc)
}
