package domain

// CreatePokemonQuery holds data to be persisted across 'pokemon' and 'pokemon_types' tables
type CreatePokemonQuery struct {
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
	TypeIDs      []int // IDs of types to insert into pokemon_types (slot 1, slot 2)
}
