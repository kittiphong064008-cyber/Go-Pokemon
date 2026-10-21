package server

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/fiber/v3"
	configs "github.com/kittiphong064008-cyber/Go-Pokemon/config"
	"github.com/kittiphong064008-cyber/Go-Pokemon/server/router"
)

type Server struct {
	app    *fiber.App
	config *configs.Configs
	db     *sql.DB
}

var (
	once   sync.Once
	server *Server
)

func NewServer(config *configs.Configs, db *sql.DB) *Server {
	once.Do(func() {
		app := fiber.New()
		server = &Server{
			app:    app,
			config: config,
			db:     db,
		}
	})
	return server
}

func (s *Server) Start() {
	api := s.app.Group("/api/v1")
	router.Pokemon(api, s.db)
	log.Printf("Server start at port = %s", s.config.App.Port)
	if err := s.app.Listen(fmt.Sprintf(":%v", s.config.App.Port)); err != nil {
		log.Fatal("Fail to start Server ", err)
	}
}
