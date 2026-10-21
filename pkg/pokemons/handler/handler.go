package handler

import (
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/dto"
	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/repository"
	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/service"
)

type handler struct {
	service service.Service
}

var val = validator.New()

func NewHandler(service service.Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreatePokemon(c fiber.Ctx) error {
	var req dto.CreatePokemonRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: err.Error()})
	}

	if err := val.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: "Invalid Body Format"})
	}

	res, err := h.service.Create(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{Message: err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ResponseOne{
		Message: "Create Success",
		Value:   *res,
	})
}

func (h *handler) GetPokemonByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: "Invalid ID"})
	}

	res, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ResponseError{Message: "Pokemon not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ResponseOne{
		Message: "GET Success",
		Value:   *res,
	})
}

func (h *handler) GetAllPokemons(c fiber.Ctx) error {
	res, err := h.service.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ResponseAll{
		Message: "GET Success",
		Value:   res,
	})
}

func (h *handler) UpdatePokemonByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: "Invalid ID"})
	}

	var req dto.PokemonUpdateRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: err.Error()})
	}

	if err := val.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: "Invalid Body Format"})
	}

	rowsAffected, err := h.service.UpdatePokemonByID(c.Context(), id, req)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ResponseError{Message: "Pokemon not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ResponseRows{
		Message:    "Update Success",
		RowsAffect: rowsAffected,
	})
}

func (h *handler) DeletePokemon(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: "Invalid ID"})
	}

	if err := h.service.Delete(c.Context(), id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ResponseError{Message: "Pokemon not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ResponseRows{
		Message:    "Delete Success",
		RowsAffect: 1,
	})
}
