package domain

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
	TypeIDs      []int
}
