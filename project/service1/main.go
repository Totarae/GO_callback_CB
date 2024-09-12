package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project/service1/config"
	"project/service1/handler"
	"project/service1/pg"
	"project/service1/repository"
	"syscall"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	http.HandleFunc("/", handler.HelloHandler)

	// Create a server with timeouts for better performance
	server := &http.Server{
		Addr:         ":9090",
		Handler:      nil, // DefaultServeMux is used
		ReadTimeout:  5 * 60,
		WriteTimeout: 10 * 60,
		IdleTimeout:  15 * 60,
	}

	log.Printf("Starting server on :9090")

	// connect&check PG db
	pgDB, err := pg.Dial()
	if err != nil {
		fmt.Errorf("pgdb.Dial failed: %w", err)
	}

	// run Postgres migrations
	if pgDB != nil {
		log.Println("Launched PostgreSQL migrations")
		if err := runPgMigrations(); err != nil {
			log.Printf("runPgMigrations failed: %w", err)
		}
	}

	objectRepo := repository.New(pgDB)

	handlerRegister := handler.LastSeenCallback{ObjectRepoInterface: objectRepo}

	http.HandleFunc("/callback", handlerRegister.CallbackHandler)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Error starting server: %s", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("closing")

	return nil

}

// runPgMigrations runs Postgres migrations
func runPgMigrations() error {
	cfg := config.Get()

	if cfg.Database.PgMigrationsPath == "" {
		return nil
	}

	log.Printf("Migrations path: %s", cfg.Database.PgMigrationsPath)

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DbName)

	log.Printf("DB path: %s", dbURL)

	m, err := migrate.New(
		cfg.Database.PgMigrationsPath,
		dbURL,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	log.Printf("Done!")
	return nil
}
