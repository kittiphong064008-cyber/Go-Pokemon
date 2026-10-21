package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/domain"
)

type mockRepository struct {
	db *sql.DB
}

func (m *mockRepository) Create(ctx context.Context, query domain.CreatePokemonQuery) (*domain.PokemonDetail, error) {
	return nil, nil
}

func (m *mockRepository) GetByID(ctx context.Context, id int) (*domain.PokemonDetail, error) {
	return nil, nil
}

func (m *mockRepository) GetAll(ctx context.Context) ([]domain.PokemonDetail, error) {
	gen1 := 1
	now := time.Now()

	results := []domain.PokemonDetail{
		{
			Pokemon: domain.Pokemon{
				ID:           1,
				Name:         "bulbasaur",
				Height:       7,
				Weight:       69,
				BaseExp:      64,
				Hp:           45,
				Attack:       49,
				Defense:      49,
				SpAtk:        65,
				SpDef:        65,
				Speed:        45,
				ImageURL:     "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/1.png",
				GenerationID: &gen1,
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			Generation: &domain.Generation{
				ID:     1,
				Name:   "generation-i",
				Region: "kanto",
			},
			Types: []domain.PokemonType{
				{PokemonID: 1, TypeID: 5, Slot: 1, TypeName: "grass"},
				{PokemonID: 1, TypeID: 8, Slot: 2, TypeName: "poison"},
			},
		},
		{
			Pokemon: domain.Pokemon{
				ID:           4,
				Name:         "charmander",
				Height:       6,
				Weight:       85,
				BaseExp:      62,
				Hp:           39,
				Attack:       52,
				Defense:      43,
				SpAtk:        60,
				SpDef:        50,
				Speed:        65,
				ImageURL:     "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/4.png",
				GenerationID: &gen1,
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			Generation: &domain.Generation{
				ID:     1,
				Name:   "generation-i",
				Region: "kanto",
			},
			Types: []domain.PokemonType{
				{PokemonID: 4, TypeID: 2, Slot: 1, TypeName: "fire"},
			},
		},
		{
			Pokemon: domain.Pokemon{
				ID:           25,
				Name:         "pikachu",
				Height:       4,
				Weight:       60,
				BaseExp:      112,
				Hp:           35,
				Attack:       55,
				Defense:      40,
				SpAtk:        50,
				SpDef:        50,
				Speed:        90,
				ImageURL:     "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/25.png",
				GenerationID: &gen1,
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			Generation: &domain.Generation{
				ID:     1,
				Name:   "generation-i",
				Region: "kanto",
			},
			Types: []domain.PokemonType{
				{PokemonID: 25, TypeID: 4, Slot: 1, TypeName: "electric"},
			},
		},
	}

	return results, nil
}

func (m *mockRepository) Update(ctx context.Context, id int, query domain.CreatePokemonQuery) (int, error) {
	return 0, nil
}

func (m *mockRepository) Delete(ctx context.Context, id int) error {
	return nil
}
