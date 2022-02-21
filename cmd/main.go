package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/darragh-downey/stanley/pkg/handlers"
)

func main() {
	log.Println("Loading server")

	// graceful shutdown
	// https://github.com/gorilla/mux#graceful-shutdown

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file - exiting: %v", err)
		os.Exit(0)
	}

	var wait time.Duration

	addr := os.Getenv("ADDR")
	port := os.Getenv("PORT")

	addr_str := fmt.Sprintf("%s:%s", addr, port)

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.JSONLinearHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         addr_str,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Println("server ready!")

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("shutting down...")
	os.Exit(0)
}
