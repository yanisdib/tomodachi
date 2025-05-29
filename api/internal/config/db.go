package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func OpenDBConnectionPool() *pgxpool.Pool {
	connectionString := buildConnectionString()
	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to reach database: %v\n", err)
	}

	fmt.Println("Connected to database")
	return pool
}

func buildConnectionString() string {
	user, err := os.ReadFile("/run/secrets/db_user")
	if err != nil {
		log.Fatalf("error occurred while reading db_user secret: %v", err)
	}

	password, err := os.ReadFile("/run/secrets/db_password")
	if err != nil {
		log.Fatalf("error occurred while reading db_password secret: %v", err)
	}

	host, err := os.ReadFile("/run/secrets/db_host")
	if err != nil {
		log.Fatalf("error occurred while reading db_host secret: %v", err)
	}

	dbName, err := os.ReadFile("/run/secrets/db_name")
	if err != nil {
		log.Fatalf("error occurred while reading db_name secret: %v", err)
	}

	port := os.Getenv("POSTGRESQL_PORT")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		string(host),
		port,
		string(user),
		string(password),
		string(dbName),
	)
}
