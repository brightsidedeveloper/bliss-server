package db

import (
	"bliss-server/genesis/queries"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func CreateDB(dsn string) (*queries.Queries, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	queries := queries.New(conn)

	log.Println("Database connected successfully!")
	return queries, nil
}
