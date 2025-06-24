package templates

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

func PostgresRepo(pkgPath string, module string, fields []Field) string {
	name := strcase.ToCamel(module)

	fieldNames := []string{}
	placeholders := []string{}
	args := []string{}
	scanFields := []string{}
	updateAssignments := []string{}
	updateArgs := []string{}
	i := 1

	for _, f := range fields {
		fieldNames = append(fieldNames, strcase.ToSnake(f.Name))
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
		args = append(args, fmt.Sprintf("%s.%s", module, strcase.ToCamel(f.Name)))
		scanFields = append(scanFields, "&obj."+strcase.ToCamel(f.Name))
		updateAssignments = append(updateAssignments, fmt.Sprintf("%s=$%d", strcase.ToSnake(f.Name), i))
		updateArgs = append(updateArgs, fmt.Sprintf("%s.%s", module, strcase.ToCamel(f.Name)))
		i++
	}

	insertFields := strings.Join(fieldNames, ", ")
	insertPlaceholders := strings.Join(placeholders, ", ")
	insertArgs := strings.Join(args, ", ")
	scanArgs := strings.Join(scanFields, ", ")
	updateSet := strings.Join(updateAssignments, ", ")
	updateArgsStr := strings.Join(updateArgs, ", ")

	return fmt.Sprintf(`package postgres

import (
    "database/sql"
    "errors"
    "fmt"
    "strings"
    "%s/internal/%s"
)

type %sRepo struct {
    db *sql.DB
}

func New(db *sql.DB) %s.Repository {
    return &%sRepo{db: db}
}

func (r *%sRepo) Create(%s *%s.%s) error {
    _, err := r.db.Exec("INSERT INTO %ss (%s) VALUES (%s)", %s)
    return err
}

func (r *%sRepo) FindAll(page, limit int, sort, order, search string) ([]%s.%s, int64, error) {
    // Build the base query
    baseQuery := "FROM %ss WHERE deleted_at IS NULL"
    countQuery := fmt.Sprintf("SELECT COUNT(*) %%s", baseQuery)
    
    args := []interface{}{}
    
    // Add search condition if provided
    if search != "" {
        searchFields := []string{%s}
        searchConditions := make([]string, len(searchFields))
        for i, field := range searchFields {
            searchConditions[i] = fmt.Sprintf("LOWER(%s::text) LIKE LOWER($1)", field)
        }
        baseQuery += fmt.Sprintf(" AND (%s)", strings.Join(searchConditions, " OR "))
        args = append(args, "%%"+search+"%%")
    }
    
    // Get total count
    var total int64
    err := r.db.QueryRow(fmt.Sprintf(countQuery, baseQuery), args...).Scan(&total)
    if err != nil {
        return nil, 0, err
    }
    
    // Add sorting
    if sort != "" {
        if order == "" {
            order = "asc"
        }
        baseQuery += fmt.Sprintf(" ORDER BY %%s %%s", sort, order)
    }
    
    // Add pagination
    offset := (page - 1) * limit
    baseQuery += fmt.Sprintf(" LIMIT %%d OFFSET %%d", limit, offset)
    
    // Execute the final query
    rows, err := r.db.Query(fmt.Sprintf("SELECT %s %%s", baseQuery), args...)
    if err != nil {
        return nil, 0, err
    }
    defer rows.Close()
    
    var result []%s.%s
    for rows.Next() {
        var obj %s.%s
        if err := rows.Scan(%s); err != nil {
            return nil, 0, err
        }
        result = append(result, obj)
    }
    
    return result, total, nil
}

func (r *%sRepo) FindByID(id string) (*%s.%s, error) {
    var obj %s.%s
    err := r.db.QueryRow("SELECT %s FROM %ss WHERE id=$1 AND deleted_at IS NULL", id).Scan(%s)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("not found")
        }
        return nil, err
    }
    return &obj, nil
}

func (r *%sRepo) Update(id string, %s *%s.%s) error {
    _, err := r.db.Exec("UPDATE %ss SET %s WHERE id=$%d AND deleted_at IS NULL", %s, id)
    return err
}

func (r *%sRepo) Delete(id string) error {
    _, err := r.db.Exec("UPDATE %ss SET deleted_at=NOW() WHERE id=$1 AND deleted_at IS NULL", id)
    return err
}
`, pkgPath, module,
		module, module, module,
		module, module, module, name, module,
		insertFields, insertPlaceholders, insertArgs,
		module, module, name, module,
		strings.Join(fieldNames, ", "),
		insertFields, module, module, name, module, name, scanArgs,
		module, module, name, module, name, insertFields, module, scanArgs,
		module, module, module, name, module, updateSet, i, updateArgsStr,
		module, module)
}
