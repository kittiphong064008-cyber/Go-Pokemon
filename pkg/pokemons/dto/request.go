package dto

type CreatePokemonRequest struct {
	ID           int    `json:"id" validate:"required,gt=0"`
	Name         string `json:"name" validate:"required,min=2,max=50"`
	Height       int    `json:"height" validate:"required,gt=0"`
	Weight       int    `json:"weight" validate:"required,gt=0"`
	BaseExp      int    `json:"base_exp" validate:"omitempty,gte=0"`
	Hp           int    `json:"hp" validate:"required,gte=1,lte=1000"`
	Attack       int    `json:"attack" validate:"required,gte=1,lte=500"`
	Defense      int    `json:"defense" validate:"required,gte=1,lte=500"`
	SpAtk        int    `json:"sp_atk" validate:"required,gte=1,lte=500"`
	SpDef        int    `json:"sp_def" validate:"required,gte=1,lte=500"`
	Speed        int    `json:"speed" validate:"required,gte=1,lte=500"`
	ImageURL     string `json:"image_url" validate:"omitempty,url"`
	GenerationID *int   `json:"generation_id" validate:"omitempty,gt=0"`
	TypeIDs      []int  `json:"type_ids" validate:"required,min=1,max=2"` // slot 1 & optional slot 2
}
