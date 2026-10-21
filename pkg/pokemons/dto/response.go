package dto

import "time"

type PokemonStats struct {
	Hp      int `json:"hp"`
	Attack  int `json:"attack"`
	Defense int `json:"defense"`
	SpAtk   int `json:"sp_atk"`
	SpDef   int `json:"sp_def"`
	Speed   int `json:"speed"`
}

type PokemonTypeResponse struct {
	Slot int    `json:"slot"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GenerationResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Region string `json:"region"`
}

type PokemonDetailResponse struct {
	ID         int                   `json:"id"`
	Name       string                `json:"name"`
	Height     int                   `json:"height"`
	Weight     int                   `json:"weight"`
	BaseExp    int                   `json:"base_exp"`
	Stats      PokemonStats          `json:"stats"`
	ImageURL   string                `json:"image_url,omitempty"`
	Generation *GenerationResponse   `json:"generation,omitempty"`
	Types      []PokemonTypeResponse `json:"types"`
	CreatedAt  time.Time             `json:"created_at"`
	UpdatedAt  time.Time             `json:"updated_at"`
}

type PokemonListItemResponse struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	ImageURL string   `json:"image_url,omitempty"`
	Types    []string `json:"types"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}
