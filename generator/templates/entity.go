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
