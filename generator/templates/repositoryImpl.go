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
		args = append(args, fmt.Sprintf("%s.%s", module, f.Name))
		scanFields = append(scanFields, "&obj."+f.Name)
		updateAssignments = append(updateAssignments, fmt.Sprintf("%s=$%d", strcase.ToSnake(f.Name), i))
		updateArgs = append(updateArgs, fmt.Sprintf("%s.%s", module, f.Name))
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

func (r *%sRepo) FindAll() ([]%s.%s, error) {
rows, err := r.db.Query("SELECT %s FROM %ss")
if err != nil {
return nil, err
}
defer rows.Close()

var result []%s.%s
for rows.Next() {
var obj %s.%s
if err := rows.Scan(%s); err != nil {
return nil, err
}
result = append(result, obj)
}
return result, nil
}

func (r *%sRepo) FindByID(id string) (*%s.%s, error) {
var obj %s.%s
err := r.db.QueryRow("SELECT %s FROM %ss WHERE id=$1", id).Scan(%s)
if err != nil {
if err == sql.ErrNoRows {
return nil, errors.New("not found")
}
return nil, err
}
return &obj, nil
}

func (r *%sRepo) Update(id string, %s *%s.%s) error {
_, err := r.db.Exec("UPDATE %ss SET %s WHERE id=$%d", %s, id)
return err
}

func (r *%sRepo) Delete(id string) error {
_, err := r.db.Exec("UPDATE %ss SET deleted_at=NOW() WHERE id=$1", id)
return err
}
`, pkgPath, module, module, module, module,
		module, module, module, name, module,
		module, insertFields, insertPlaceholders, insertArgs,
		module, module, module, insertFields, module,
		module, module, name, module, scanArgs,
		module, module, module, name, module, insertFields, module, scanArgs,
		module, module, module, name, module, module, updateSet, i, updateArgsStr,
		module, module, module)
}
