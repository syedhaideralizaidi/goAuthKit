package main

import (
	"context"
	"github.com/syedhaideralizaidi/authkit/internal/database"
	"github.com/syedhaideralizaidi/authkit/internal/route"
	"log"
)

func main() {
	database.ConnectDB()
	defer database.Conn.Close(context.Background())

	r := route.SetupRouter()

	err := r.Run(":8080")
	log.Fatal(err)
}
