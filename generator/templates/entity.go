package templates

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

func Entity(module string, fields []Field) string {
	name := strcase.ToCamel(module)
	structFields := ""
	for _, f := range fields {
		tag := ""
		if f.Tag != "" {
			tag = " " + f.Tag
		}
		structFields += fmt.Sprintf("%s %s%s\n", strcase.ToCamel(f.Name), f.Type, tag)
	}
	return fmt.Sprintf(`package %s

type %s struct {
%s}`, module, name, structFields)
}
