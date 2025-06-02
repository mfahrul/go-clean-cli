package templates

import "fmt"

func DTO(module string, fields []Field) string {
	createFields := ""
	updateFields := ""
	responseFields := ""

	for _, f := range fields {
		tag := ""
		if f.Tag != "" {
			tag = " " + f.Tag
		}
		createFields += fmt.Sprintf("%s %s%s\n", f.Name, f.Type, tag)
		updateFields += fmt.Sprintf("%s %s%s\n", f.Name, f.Type, tag)
		responseFields += fmt.Sprintf("%s %s%s\n", f.Name, f.Type, tag)
	}

	return fmt.Sprintf(`package %s

type CreateDTO struct {
%s}

type UpdateDTO struct {
%s}

type ResponseDTO struct {
%s
CreatedAt *time.Time `+"`json:\"created_at,omitempty\"`"+`
UpdatedAt *time.Time `+"`json:\"updated_at,omitempty\"`"+`
DeletedAt *time.Time `+"`json:\"deleted_at,omitempty\"`"+`
}
`, module, createFields, updateFields, responseFields)
}
