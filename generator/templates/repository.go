package templates

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

func Repository(module string, fields []Field) string {
	name := strcase.ToCamel(module)
	return fmt.Sprintf(`package %s

type Repository interface {
Create(%s *%s) error
FindAll(page, limit int, sort, order, search string) ([]%s, int64, error)
FindByID(id string) (*%s, error)
Update(id string, %s *%s) error
Delete(id string) error
}`, module, module, name, name, name, module, name)
}
