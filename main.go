package main

import (
	"log"

	configs "github.com/kittiphong064008-cyber/Go-Pokemon/config"
	"github.com/kittiphong064008-cyber/Go-Pokemon/databases"
	"github.com/kittiphong064008-cyber/Go-Pokemon/server"
)

func main() {
	cfg := configs.Load()
	db, err := databases.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := server.NewServer(&cfg, db)
	app.Start()
}
