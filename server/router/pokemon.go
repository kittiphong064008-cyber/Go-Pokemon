package router

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/handler"
	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/repository"
	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/service"
)

func Pokemon(router fiber.Router, db *sql.DB) {
	pokemonRepo := repository.NewRepository(db)
	pokemonService := service.NewService(pokemonRepo)
	pokemonHandler := handler.NewHandler(pokemonService)

	group := router.Group("/pokemon")
	group.Post("/", pokemonHandler.CreatePokemon)
	group.Get("/", pokemonHandler.GetAllPokemons)
	group.Get("/:id", pokemonHandler.GetPokemonByID)
	group.Patch("/:id", pokemonHandler.UpdatePokemonByID)
	group.Delete("/:id", pokemonHandler.DeletePokemon)
}
