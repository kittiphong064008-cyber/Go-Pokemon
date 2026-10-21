package domain

import "time"

// Pokemon represents the core 'pokemon' table entity
type Pokemon struct {
	ID           int
	Name         string
	Height       int
	Weight       int
	BaseExp      int
	Hp           int
	Attack       int
	Defense      int
	SpAtk        int
	SpDef        int
	Speed        int
	ImageURL     string
	GenerationID *int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Generation represents the 'generations' lookup table
type Generation struct {
	ID     int
	Name   string
	Region string
}

// Type represents the 'types' lookup table
type Type struct {
	ID   int
	Name string
}

// PokemonType represents relation in 'pokemon_types'
type PokemonType struct {
	PokemonID int
	TypeID    int
	Slot      int
	TypeName  string
}

// PokemonDetail combines Pokemon with its relations (Types, Generation)
type PokemonDetail struct {
	Pokemon    Pokemon
	Generation *Generation
	Types      []PokemonType
}
