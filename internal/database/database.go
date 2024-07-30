package database

import (
	"context"
	"log"

	db "github.com/syedhaideralizaidi/authkit/internal/database/sqlc"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn
var Queries *db.Queries

func ConnectDB() {
	var err error
	Conn, err = pgx.Connect(context.Background(), "postgres://root:secret@localhost:5432/authKit?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	Queries = db.New(Conn) // Initialize Queries with the connection
}
