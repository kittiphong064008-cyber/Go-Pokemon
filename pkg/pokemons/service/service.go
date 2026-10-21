package service

import (
	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/domain"
	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/dto"
	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/repository"
)

type Service interface {
	Create(req dto.CreatePokemonRequest) (*dto.PokemonDetailResponse, error)
	GetByID(id int) (*dto.PokemonDetailResponse, error)
	GetAll() ([]dto.PokemonDetailResponse, error)
	Delete(id int) error
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(req dto.CreatePokemonRequest) (*dto.PokemonDetailResponse, error) {
	query := domain.CreatePokemonQuery{
		ID:           req.ID,
		Name:         req.Name,
		Height:       req.Height,
		Weight:       req.Weight,
		BaseExp:      req.BaseExp,
		Hp:           req.Hp,
		Attack:       req.Attack,
		Defense:      req.Defense,
		SpAtk:        req.SpAtk,
		SpDef:        req.SpDef,
		Speed:        req.Speed,
		ImageURL:     req.ImageURL,
		GenerationID: req.GenerationID,
		TypeIDs:      req.TypeIDs,
	}

	detail, err := s.repo.Create(query)
	if err != nil {
		return nil, err
	}

	return toDetailResponse(detail), nil
}

func (s *service) GetByID(id int) (*dto.PokemonDetailResponse, error) {
	detail, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return toDetailResponse(detail), nil
}

func (s *service) GetAll() ([]dto.PokemonDetailResponse, error) {
	details, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.PokemonDetailResponse, 0, len(details))
	for _, d := range details {
		responses = append(responses, *toDetailResponse(&d))
	}

	return responses, nil
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}

func toDetailResponse(d *domain.PokemonDetail) *dto.PokemonDetailResponse {
	typesResp := make([]dto.PokemonTypeResponse, 0, len(d.Types))
	for _, t := range d.Types {
		typesResp = append(typesResp, dto.PokemonTypeResponse{
			Slot: t.Slot,
			ID:   t.TypeID,
			Name: t.TypeName,
		})
	}

	var genResp *dto.GenerationResponse
	if d.Generation != nil {
		genResp = &dto.GenerationResponse{
			ID:     d.Generation.ID,
			Name:   d.Generation.Name,
			Region: d.Generation.Region,
		}
	}

	return &dto.PokemonDetailResponse{
		ID:      d.Pokemon.ID,
		Name:    d.Pokemon.Name,
		Height:  d.Pokemon.Height,
		Weight:  d.Pokemon.Weight,
		BaseExp: d.Pokemon.BaseExp,
		Stats: dto.PokemonStats{
			Hp:      d.Pokemon.Hp,
			Attack:  d.Pokemon.Attack,
			Defense: d.Pokemon.Defense,
			SpAtk:   d.Pokemon.SpAtk,
			SpDef:   d.Pokemon.SpDef,
			Speed:   d.Pokemon.Speed,
		},
		ImageURL:   d.Pokemon.ImageURL,
		Generation: genResp,
		Types:      typesResp,
		CreatedAt:  d.Pokemon.CreatedAt,
		UpdatedAt:  d.Pokemon.UpdatedAt,
	}
}
