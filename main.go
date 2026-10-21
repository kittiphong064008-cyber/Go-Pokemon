package main

import (
	"fmt"
	"log"

	"github.com/kittiphong064008-cyber/Go-Pokemon/config"
	"github.com/kittiphong064008-cyber/Go-Pokemon/databases"
	"github.com/kittiphong064008-cyber/Go-Pokemon/server"
)

func main() {
	cfg := config.Load()

	db, err := databases.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer db.Close()

	app := server.NewServer(db, cfg)

	addr := fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port)
	log.Printf("Starting server on http://%s", addr)
	log.Fatal(app.Listen(addr))
}
