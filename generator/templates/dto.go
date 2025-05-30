package templates

import "fmt"

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
