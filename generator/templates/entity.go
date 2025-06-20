package templates

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

func Entity(module string, fields []Field) string {
	name := strcase.ToCamel(module)
	structFields := ""
	for _, f := range fields {
		jsonTag := fmt.Sprintf("`json:\"%s\"`", strcase.ToSnake(f.Name))
		structFields += fmt.Sprintf("%s %s %s\n\t", strcase.ToCamel(f.Name), f.Type, jsonTag)
	}
	return fmt.Sprintf(`package %s

import "time"

type %s struct {
	%sCreatedAt *time.Time `+"`json:\"created_at,omitempty\"`"+`
	UpdatedAt *time.Time `+"`json:\"updated_at,omitempty\"`"+`
	DeletedAt *time.Time `+"`json:\"deleted_at,omitempty\"`"+`
}`, module, name, structFields)
}
