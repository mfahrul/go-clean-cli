package templates

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

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
