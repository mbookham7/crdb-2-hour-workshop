package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	log.SetFlags(0)

	connectionString, ok := os.LookupEnv("CONNECTION_STRING")
	if !ok {
		log.Fatalf("missing CONNECTION_STRING env var")
	}

	db, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query(context.Background(), "SELECT id, full_name FROM member")
	if err != nil {
		log.Fatalf("error making query: %v", err)
	}

	var id, name string
	for rows.Next() {
		if err = rows.Scan(&id, &name); err != nil {
			log.Fatalf("error scanning row: %v", err)
		}

		log.Printf("%s: %s", id, name)
	}
}
