package service

import (
	"context"

	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/domain"
	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/dto"
	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/repository"
)

type Service interface {
	Create(ctx context.Context, req dto.CreatePokemonRequest) (*dto.PokemonDetailResponse, error)
	GetByID(ctx context.Context, id int) (*dto.PokemonDetailResponse, error)
	GetAll(ctx context.Context) ([]dto.PokemonDetailResponse, error)
	UpdatePokemonByID(ctx context.Context, id int, req dto.PokemonUpdateRequest) (int, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, req dto.CreatePokemonRequest) (*dto.PokemonDetailResponse, error) {
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

	detail, err := s.repo.Create(ctx, query)
	if err != nil {
		return nil, err
	}

	return toDetailResponse(detail), nil
}

func (s *service) GetByID(ctx context.Context, id int) (*dto.PokemonDetailResponse, error) {
	detail, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return toDetailResponse(detail), nil
}

func (s *service) GetAll(ctx context.Context) ([]dto.PokemonDetailResponse, error) {
	details, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.PokemonDetailResponse, 0, len(details))
	for _, d := range details {
		responses = append(responses, *toDetailResponse(&d))
	}

	return responses, nil
}

func (s *service) UpdatePokemonByID(ctx context.Context, id int, req dto.PokemonUpdateRequest) (int, error) {
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return 0, err
	}

	if req.Name != nil {
		existing.Pokemon.Name = *req.Name
	}
	if req.Height != nil {
		existing.Pokemon.Height = *req.Height
	}
	if req.Weight != nil {
		existing.Pokemon.Weight = *req.Weight
	}
	if req.BaseExp != nil {
		existing.Pokemon.BaseExp = *req.BaseExp
	}
	if req.Hp != nil {
		existing.Pokemon.Hp = *req.Hp
	}
	if req.Attack != nil {
		existing.Pokemon.Attack = *req.Attack
	}
	if req.Defense != nil {
		existing.Pokemon.Defense = *req.Defense
	}
	if req.SpAtk != nil {
		existing.Pokemon.SpAtk = *req.SpAtk
	}
	if req.SpDef != nil {
		existing.Pokemon.SpDef = *req.SpDef
	}
	if req.Speed != nil {
		existing.Pokemon.Speed = *req.Speed
	}
	if req.ImageURL != nil {
		existing.Pokemon.ImageURL = *req.ImageURL
	}
	if req.GenerationID != nil {
		existing.Pokemon.GenerationID = req.GenerationID
	}

	var typeIDs []int
	if len(req.TypeIDs) > 0 {
		typeIDs = req.TypeIDs
	} else {
		for _, t := range existing.Types {
			typeIDs = append(typeIDs, t.TypeID)
		}
	}

	query := domain.CreatePokemonQuery{
		ID:           existing.Pokemon.ID,
		Name:         existing.Pokemon.Name,
		Height:       existing.Pokemon.Height,
		Weight:       existing.Pokemon.Weight,
		BaseExp:      existing.Pokemon.BaseExp,
		Hp:           existing.Pokemon.Hp,
		Attack:       existing.Pokemon.Attack,
		Defense:      existing.Pokemon.Defense,
		SpAtk:        existing.Pokemon.SpAtk,
		SpDef:        existing.Pokemon.SpDef,
		Speed:        existing.Pokemon.Speed,
		ImageURL:     existing.Pokemon.ImageURL,
		GenerationID: existing.Pokemon.GenerationID,
		TypeIDs:      typeIDs,
	}

	return s.repo.Update(ctx, id, query)
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
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
