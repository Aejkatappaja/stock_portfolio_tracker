package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type dbConn struct {
	DB *sqlx.DB
}

func initDB() (*dbConn, error) {
	log.Printf("Initializing postgres database\n")

	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	pgUser := os.Getenv("PG_USER")
	pgPassword := os.Getenv("PG_PASSWORD")
	pgDB := os.Getenv("PG_DB")
	pgSSL := os.Getenv("PG_SSL")

	pgConnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", pgHost, pgPort, pgUser, pgPassword, pgDB, pgSSL)

	var db *sqlx.DB
	var err error
	for retries := 5; retries > 0; retries-- {
		log.Printf("Connecting to postgres\n")
		db, err = sqlx.Open("postgres", pgConnString)
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Println("Successfully connected to the database.")
				return &dbConn{DB: db}, nil
			}
		}

		log.Printf("Error connecting to db: %v. Retrying in 2 seconds...\n", err)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("could not connect to the database after several attempts: %w", err)
}

func (d *dbConn) close() error {
	if err := d.DB.Close(); err != nil {
		return fmt.Errorf("error closing Postgresql: %w", err)
	}
	return nil
}
