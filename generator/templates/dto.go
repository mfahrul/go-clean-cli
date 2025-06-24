package templates

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

func DTO(module string, fields []Field) string {
	createFields := ""
	updateFields := ""
	responseFields := ""

	for _, f := range fields {
		jsonTag := fmt.Sprintf("`json:\"%s\"`", strcase.ToSnake(f.Name))
		createFields += fmt.Sprintf("%s %s %s\n\t", strcase.ToCamel(f.Name), f.Type, jsonTag)
		updateFields += fmt.Sprintf("%s *%s %s\n\t", strcase.ToCamel(f.Name), f.Type, jsonTag)
		responseFields += fmt.Sprintf("%s %s %s\n\t", strcase.ToCamel(f.Name), f.Type, jsonTag)
	}

	return fmt.Sprintf(`package %s
	import "time"

type CreateDTO struct {
	Id string `+"`json:\"id\"`"+`
	%s
}

type UpdateDTO struct {
	%s
}

type ResponseDTO struct {
	Id string `+"`json:\"id\"`"+`
	%sCreatedAt *time.Time `+"`json:\"created_at,omitempty\"`"+`
	UpdatedAt *time.Time `+"`json:\"updated_at,omitempty\"`"+`
	DeletedAt *time.Time `+"`json:\"deleted_at,omitempty\"`"+`
}
`, module, createFields, updateFields, responseFields)
}
