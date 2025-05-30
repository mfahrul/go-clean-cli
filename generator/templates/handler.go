package templates

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

func Handler(pkgPath string, module string) string {
	name := strcase.ToCamel(module)
	return fmt.Sprintf(`package http

import (
	"github.com/gofiber/fiber/v2"
	"%s/internal/%s"
)

type Handler struct {
	usecase %s.Usecase
}

func NewHandler(uc %s.Usecase) *Handler {
	return &Handler{usecase: uc}
}

func (h *Handler) Create%s(c *fiber.Ctx) error {
	var dto %s.CreateDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}
	err := h.usecase.Create(dto)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusCreated).JSON("Created")
}

func (h *Handler) Get%ss(c *fiber.Ctx) error {
	results, err := h.usecase.FindAll()
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(results)
}

func (h *Handler) Get%s(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.usecase.FindByID(id)
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(result)
}

func (h *Handler) Update%s(c *fiber.Ctx) error {
	id := c.Params("id")
	var dto %s.UpdateDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.usecase.Update(id, dto)
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(result)
}

func (h *Handler) Delete%s(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.usecase.Delete(id); err != nil {
		return fiber.ErrNotFound
	}
	return c.SendStatus(fiber.StatusNoContent)
}`, pkgPath, module, module, module, name,
		module, name, name, name, module,
		name)
}
