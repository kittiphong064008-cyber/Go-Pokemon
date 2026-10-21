package domain

import "time"

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

type Generation struct {
	ID     int
	Name   string
	Region string
}

type Type struct {
	ID   int
	Name string
}

type PokemonType struct {
	PokemonID int
	TypeID    int
	Slot      int
	TypeName  string
}

type PokemonDetail struct {
	Pokemon    Pokemon
	Generation *Generation
	Types      []PokemonType
}
