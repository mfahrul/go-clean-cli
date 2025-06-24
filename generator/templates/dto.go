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

type PaginationQuery struct {
    Page   int    %s
    Limit  int    %s
    Sort   string %s
    Order  string %s
    Search string %s
}

type PaginationMeta struct {
    CurrentPage  int    %s
    TotalPages   int    %s
    TotalItems   int64  %s
    ItemsPerPage int    %s
    HasNext      bool   %s
    HasPrev      bool   %s
    NextPage     int    %s
    PrevPage     int    %s
    NextPageURL  string %s
    PrevPageURL  string %s
}

type PaginationResponse struct {
    Data []ResponseDTO  %s
    Meta PaginationMeta %s
}

type CreateDTO struct {
    Id string %s
    %s
}

type UpdateDTO struct {
    %s
}

type ResponseDTO struct {
    Id string %s
    %sCreatedAt *time.Time %s
    UpdatedAt *time.Time %s
    DeletedAt *time.Time %s
}
`,
		module,
		"`query:\"page\" json:\"page\"`",
		"`query:\"limit\" json:\"limit\"`",
		"`query:\"sort\" json:\"sort\"`",
		"`query:\"order\" json:\"order\"`",
		"`query:\"search\" json:\"search\"`",
		"`json:\"current_page\"`",
		"`json:\"total_pages\"`",
		"`json:\"total_items\"`",
		"`json:\"items_per_page\"`",
		"`json:\"has_next\"`",
		"`json:\"has_prev\"`",
		"`json:\"next_page\"`",
		"`json:\"prev_page\"`",
		"`json:\"next_page_url\"`",
		"`json:\"prev_page_url\"`",
		"`json:\"data\"`",
		"`json:\"meta\"`",
		"`json:\"id\"`",
		createFields,
		updateFields,
		"`json:\"id\"`",
		responseFields,
		"`json:\"created_at,omitempty\"`",
		"`json:\"updated_at,omitempty\"`",
		"`json:\"deleted_at,omitempty\"`")
}
