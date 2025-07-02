package templates

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

func Handler(pkgPath string, module string) string {
	name := strcase.ToCamel(module)
	return fmt.Sprintf(`package http

import (
    "fmt"
    "regexp"
    "strings"
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

func (h *Handler) buildPaginationURL(c *fiber.Ctx, page int) string {
	if page == 0 {
		return ""
	}
	baseURL := c.BaseURL()
	path := c.Path()
	queries := c.Queries()
	var queryParts []string
	for key, value := range queries {
		if key == "page" {
			// Skip the old page parameter as we'll add the new one
			continue
		}
		queryParts = append(queryParts, fmt.Sprintf("%%s=%%s", key, value))
	}
	// Add the new page parameter
	queryParts = append(queryParts, fmt.Sprintf("page=%%d", page))
	query := strings.Join(queryParts, "&")
	if len(queryParts) > 0 {
		return fmt.Sprintf("%%s%%s?%%s", baseURL, path, query)
	}
	return fmt.Sprintf("%%s%%s?page=%%d", baseURL, path, page)
}

func (h *Handler) Get%ss(c *fiber.Ctx) error {
    query := new(%s.PaginationQuery)
    if err := c.QueryParser(query); err != nil {
        return fiber.ErrBadRequest
    }

    // Set defaults if not provided
    if query.Page < 1 {
        query.Page = 1
    }
    if query.Limit < 1 {
        query.Limit = 10
    }

    // Validate sort and order
    validSortFields := map[string]bool{"created_at": true, "updated_at": true, "id": true}
    if query.Sort != "" && !validSortFields[query.Sort] {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid sort field")
    }

    validOrders := map[string]bool{"asc": true, "desc": true}
    if query.Order != "" && !validOrders[query.Order] {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid order value")
    }

    results, err := h.usecase.FindAll(query)
    if err != nil {
        return fiber.ErrInternalServerError
    }

    // Add pagination URLs
    if results.Meta.NextPage > 0 {
        results.Meta.NextPageURL = h.buildPaginationURL(c, results.Meta.NextPage)
    }
    if results.Meta.PrevPage > 0 {
        results.Meta.PrevPageURL = h.buildPaginationURL(c, results.Meta.PrevPage)
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
		module, name, module, name, name, module, name)
}
