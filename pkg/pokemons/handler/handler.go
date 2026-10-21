package handler

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/dto"
	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/repository"
	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/service"
)

type Handler struct {
	service  service.Service
	validate *validator.Validate
}

func NewHandler(service service.Service, validate *validator.Validate) *Handler {
	return &Handler{
		service:  service,
		validate: validate,
	}
}

func (h *Handler) CreatePokemon(c fiber.Ctx) error {
	var req dto.CreatePokemonRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.APIResponse{
			Success: false,
			Message: "invalid request body format",
			Errors:  err.Error(),
		})
	}

	if err := h.validate.Struct(req); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("field '%s' failed on validation rule '%s'", err.Field(), err.Tag()))
		}
		return c.Status(fiber.StatusBadRequest).JSON(dto.APIResponse{
			Success: false,
			Message: "validation error",
			Errors:  validationErrors,
		})
	}

	res, err := h.service.Create(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.APIResponse{
			Success: false,
			Message: "failed to create pokemon",
			Errors:  err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.APIResponse{
		Success: true,
		Message: "pokemon created successfully",
		Data:    res,
	})
}

func (h *Handler) GetPokemonByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.APIResponse{
			Success: false,
			Message: "invalid pokemon ID",
		})
	}

	res, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.APIResponse{
				Success: false,
				Message: "pokemon not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.APIResponse{
			Success: false,
			Message: "failed to get pokemon",
			Errors:  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.APIResponse{
		Success: true,
		Message: "pokemon retrieved successfully",
		Data:    res,
	})
}

func (h *Handler) GetAllPokemons(c fiber.Ctx) error {
	res, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.APIResponse{
			Success: false,
			Message: "failed to retrieve pokemons",
			Errors:  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.APIResponse{
		Success: true,
		Message: "pokemons retrieved successfully",
		Data:    res,
	})
}

func (h *Handler) DeletePokemon(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.APIResponse{
			Success: false,
			Message: "invalid pokemon ID",
		})
	}

	if err := h.service.Delete(id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.APIResponse{
				Success: false,
				Message: "pokemon not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.APIResponse{
			Success: false,
			Message: "failed to delete pokemon",
			Errors:  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.APIResponse{
		Success: true,
		Message: "pokemon deleted successfully",
	})
}
