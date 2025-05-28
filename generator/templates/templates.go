package templates

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

func Entity(module string) string {
	name := strcase.ToCamel(module)
	return fmt.Sprintf(`package %s

import "time"

type %s struct {
	ID    string  `+"`json:\"id\"`"+`
	Title string `+"`json:\"title\"`"+`
	Body  string `+"`json:\"body\"`"+`
	CreatedAt *time.Time `+"`json:\"created_at,omitempty\"`"+`
	UpdatedAt *time.Time `+"`json:\"updated_at,omitempty\"`"+`
	DeletedAt *time.Time `+"`json:\"deleted_at,omitempty\"`"+`
}`, module, name)
}

func DTO(module string) string {
	return fmt.Sprintf(`package %s

import "time"

type CreateDTO struct {
	Title string `+"`json:\"title\"`"+`
	Body  string `+"`json:\"body\"`"+`
}

type UpdateDTO struct {
	Title string `+"`json:\"title\"`"+`
	Body  string `+"`json:\"body\"`"+`
}
	
type ResponseDTO struct {
	ID    string  `+"`json:\"id\"`"+`
	Title string `+"`json:\"title\"`"+`
	Body  string `+"`json:\"body\"`"+`
	CreatedAt *time.Time `+"`json:\"created_at,omitempty\"`"+`
	UpdatedAt *time.Time `+"`json:\"updated_at,omitempty\"`"+`
	DeletedAt *time.Time `+"`json:\"deleted_at,omitempty\"`"+`
}`, module)
}

func Repository(module string) string {
	name := strcase.ToCamel(module)
	return fmt.Sprintf(`package %s

type Repository interface {
	Create(%s *%s) error
	FindAll() ([]%s, error)
	FindByID(id string) (*%s, error)
	Update(id string, %s *%s) error
	Delete(id string) error
}`, module, module, name, name, name,
		module, name)
}

func Usecase(module string) string {
	return fmt.Sprintf(`package %s

type Usecase interface {
	Create(dto CreateDTO) (error)
	FindAll() ([]ResponseDTO, error)
	FindByID(id string) (*ResponseDTO, error)
	Update(id string, dto UpdateDTO) (*ResponseDTO, error)
	Delete(id string) error
}`, module)
}

func UsecaseImpl(pkgPath string, module string) string {
	name := strcase.ToCamel(module)
	return fmt.Sprintf(`package usecase

import (
	"fmt"
	"%s/internal/%s"
)
type %sUsecase struct {
	repo %s.Repository
}

func New(repo %s.Repository) %s.Usecase {
	return &%sUsecase{repo: repo}
}

func (uc *%sUsecase) Create(dto %s.CreateDTO) (error) {
	obj := &%s.%s{
		Title: dto.Title,
		Body:  dto.Body,
	}
	return uc.repo.Create(obj)
}

func (uc *%sUsecase) FindAll() ([]%s.ResponseDTO, error) {
	%ss, err := uc.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var responseDTOs []%s.ResponseDTO
	for _, obj := range %ss {
		respDTO := &%s.ResponseDTO{
			ID:    obj.ID,
			Title: obj.Title,
			Body:  obj.Body,
			CreatedAt: obj.CreatedAt,
			UpdatedAt: obj.UpdatedAt,
			DeletedAt: obj.DeletedAt,
		}
		responseDTOs = append(responseDTOs, *respDTO)
	}
	return responseDTOs, nil
}

func (uc *%sUsecase) FindByID(id string) (*%s.ResponseDTO, error) {
	obj, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	result := &%s.ResponseDTO{
		ID:    obj.ID,
		Title: obj.Title,
		Body:  obj.Body,
		CreatedAt: obj.CreatedAt,
		UpdatedAt: obj.UpdatedAt,
		DeletedAt: obj.DeletedAt,
	}
	return result, nil
}

func (uc *%sUsecase) Update(id string, dto %s.UpdateDTO) (*%s.ResponseDTO, error) {
	obj, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	var changed = false
	if obj == nil {
		return nil, fmt.Errorf("object with id $s not found", id)
	}
	if obj.Title != dto.Title && dto.Title != "" {
		if !changed {
			changed = true
		}
		obj.Title = dto.Title
	}
	if obj.Body != dto.Body && dto.Body != "" {
		if !changed {
			changed = true
		}
		obj.Body = dto.Body
	}
	if !changed {
		return nil, fmt.Errorf("no changes detected")
	}
	err = uc.repo.Update(id, obj)
	if err != nil {
		return nil, err
	}
	result := &%s.ResponseDTO{
		ID:    obj.ID,
		Title: obj.Title,
		Body:  obj.Body,
		CreatedAt: obj.CreatedAt,
		UpdatedAt: obj.UpdatedAt,
		DeletedAt: obj.DeletedAt,
	}
	return result, nil
}

func (uc *%sUsecase) Delete(id string) error {
	return uc.repo.Delete(id)
}`, pkgPath, module, module, module, module,
		module, module, module, module, module,
		name, module, module, module, module,
		module, module, module, module, module,
		module, module, module, module, module)
}

func PostgresRepo(pkgPath string, module string) string {
	name := strcase.ToCamel(module)
	return fmt.Sprintf(`package postgres

import (
	"database/sql"
	"errors"
	"%s/internal/%s"
)

type %sRepo struct {
	db *sql.DB
}

func New(db *sql.DB) %s.Repository {
	return &%sRepo{db: db}
}

func (r *%sRepo) Create(%s *%s.%s) error {
	_, err := r.db.Exec("INSERT INTO %ss (title, body) VALUES ($1, $2)", %s.Title, %s.Body)
	return err
}

func (r *%sRepo) FindAll() ([]%s.%s, error) {
	rows, err := r.db.Query("SELECT id, title, body, created_at, updated_at, deleted_at FROM %ss")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []%s.%s
	for rows.Next() {
		var obj %s.%s
		if err := rows.Scan(&obj.ID, &obj.Title, &obj.Body, &obj.CreatedAt, &obj.UpdatedAt, &obj.DeletedAt); err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

func (r *%sRepo) FindByID(id string) (*%s.%s, error) {
	var obj %s.%s
	err := r.db.QueryRow("SELECT id, title, body, created_at, updated_at, deleted_at FROM %ss WHERE id=$1", id).Scan(&obj.ID, &obj.Title, &obj.Body, &obj.CreatedAt, &obj.UpdatedAt, &obj.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("not found")
		}
		return nil, err
	}
	return &obj, nil
}

func (r *%sRepo) Update(id string, %s *%s.%s) error {
	_, err := r.db.Exec("UPDATE %ss SET title=$1, body=$2, updated_at=NOW() WHERE id=$3", %s.Title, %s.Body, id)
	return err
}

func (r *%sRepo) Delete(id string) error {
	// Soft delete
	_, err := r.db.Exec("UPDATE %ss SET deleted_at=NOW() WHERE id=$1", id)
	// Hard delete
	// Delete the user from the database
	// This will remove the user permanently
	// and not just mark it as deleted.
	// _, err := r.db.Exec("DELETE FROM %ss WHERE id=$1", id)	
	return err
}`, pkgPath, module, module, module, module,
		module, module, module, name, module,
		module, module, module, module, name,
		module, module, name, module, name,
		module, module, name, module, name,
		module, module, module, module, name,
		module, module, module, module, module,
		module)
}

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

func Routes(pkgPath string, module string) string {
	name := strcase.ToCamel(module)
	return fmt.Sprintf(`package http

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"%s/core/modules"
	"%s/internal/%s/repository/postgres"
	"%s/internal/%s/usecase"	
)

func init() {
	// Register the module
	module := &modules.Module{
		Route: RegisterRoutes,
	}
	modules.RegisterModules.Add(module)
}

func RegisterRoutes(r fiber.Router, db *sql.DB) {
	repo := postgres.New(db)
	uc := usecase.New(repo)
	h := NewHandler(uc)
	group := r.Group("/%ss")
	group.Post("/", h.Create%s)
	group.Get("/", h.Get%ss)
	group.Get("/:id", h.Get%s)
	group.Put("/:id", h.Update%s)
	group.Delete("/:id", h.Delete%s)
}`, pkgPath, pkgPath, module, pkgPath, module,
		module, name, name, name, name,
		name)
}

func UnitTest(pkgPath string, module string) string {
	return fmt.Sprintf(`package test

import (
	"testing"
//	"%s/internal/%s"
)

func TestCreate(t *testing.T) {
	// TODO: implement unit test for Create
}`, pkgPath, module)
}

func IntegrationTest(pkgPath string, module string) string {
	return fmt.Sprintf(`package test

import (
	"testing"
//	"database/sql"
	_ "github.com/lib/pq"
//	"%s/internal/%s"
)

func TestIntegrationCreate(t *testing.T) {
	// TODO: implement integration test for Create with real DB
}`, pkgPath, module)
}
