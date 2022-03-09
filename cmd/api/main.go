package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/andrewchurchill/go-tutorial/models"
	_ "github.com/lib/pq"
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
	db   struct {
		dsn string
	}
}

type application struct {
	config configuration
	logger *log.Logger
	models models.Models
}

func main() {
	// Parse configuration from the command line
	var cfg configuration
	flag.IntVar(&cfg.port, "port", 4000, "The server will listen to this port.")
	flag.StringVar(&cfg.env, "env", "development", "The application environment (development|production).")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://postgres:password@localhost/go_movies?sslmode=disable", "Postgres connection string")
	flag.Parse()

	// Set up logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Set up database
	db, err := openDb(cfg)
	if err != nil {
		logger.Fatal(err)
		return
	}

	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
		models: models.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("Starting server on port %d...\n", cfg.port)

	err = srv.ListenAndServe()
	if err != nil {
		app.logger.Println(err)
	}
}

func openDb(cfg configuration) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
