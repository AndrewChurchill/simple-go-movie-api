package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type configuration struct {
	port int
	env  string
}

type application struct {
	config configuration
	logger *log.Logger
}

func main() {
	// Parse configuration from the command line
	var cfg configuration
	flag.IntVar(&cfg.port, "port", 4000, "The server will listen to this port.")
	flag.StringVar(&cfg.env, "env", "development", "The application environment (development|production).")
	flag.Parse()

	// Set up logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("Starting server on port %d...\n", cfg.port)

	err := srv.ListenAndServe()
	if err != nil {
		app.logger.Println(err)
	}
}
