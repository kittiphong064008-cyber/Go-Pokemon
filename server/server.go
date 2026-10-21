package server

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/kittiphong064008-cyber/Go-Pokemon/config"
	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/handler"
	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/repository"
	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/service"
)

func NewServer(db *sql.DB, cfg *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Go-Pokemon API v1.0",
	})

	// Validator instance
	validate := validator.New()

	// Dependency Injection: Repository -> Service -> Handler
	pokemonRepo := repository.NewRepository(db)
	pokemonService := service.NewService(pokemonRepo)
	pokemonHandler := handler.NewHandler(pokemonService, validate)

	// Health Check
	app.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "ok",
			"message": "Go-Pokemon API is running smoothly!",
		})
	})

	// API Routes Group
	api := app.Group("/api/v1")
	pokemons := api.Group("/pokemons")

	pokemons.Post("/", pokemonHandler.CreatePokemon)
	pokemons.Get("/", pokemonHandler.GetAllPokemons)
	pokemons.Get("/:id", pokemonHandler.GetPokemonByID)
	pokemons.Delete("/:id", pokemonHandler.DeletePokemon)

	return app
}
