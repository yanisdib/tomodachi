package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Creates and uses a connection pool to the database
func OpenDBConnectionPool() *pgxpool.Pool {
	connectionURL := createConnectionURL()
	dbPool, err := pgxpool.New(context.Background(), connectionURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	err = dbPool.Ping(context.TODO())
	if err != nil {
		log.Fatalf("Unable to reach database: %v", err)
	}

	fmt.Println("Connected to database")
	return dbPool
}

// Creates a database connection URL based on Docker secrets
func createConnectionURL() string {
	user, err := os.ReadFile("/run/secrets/db_user")
	if err != nil {
		log.Fatalf("Error occured while reading db_user secret: %v", err)
	}

	password, err := os.ReadFile("/run/secrets/db_password")
	if err != nil {
		log.Fatalf("Error occured while reading db_password secret: %v", err)
	}

	host, err := os.ReadFile("/run/secrets/db_host")
	if err != nil {
		log.Fatalf("Error occured while reading db_host secret: %v", err)
	}

	dbname, err := os.ReadFile("/run/secrets/db_name")
	if err != nil {
		log.Fatalf("Error occured while reading db_name secret: %v", err)
	}

	port := os.Getenv("POSTGRESQL_PORT")

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)
}
