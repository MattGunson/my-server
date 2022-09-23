package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type Store struct {
}

func PostRequest(req Request) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	conn.Exec(context.Background(),
		"INSERT INTO requests (url, headers, body) VALUES ($1, $2, $3)",
		req.Url, req.Headers, req.Body)

	fmt.Println(greeting)
	fmt.Println(req.Url)
}
