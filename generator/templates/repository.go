package templates

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

func Repository(module string) string {
	name := strcase.ToCamel(module)
	return fmt.Sprintf(`package %s

type Repository interface {
	Create(%s *%s) error
	FindAll() ([]%s, error)
	FindByID(id string) (*%s, error)
	Update(id string, %s *%s) error
	Delete(id string) error
}`, module, module, name, name, name,
		module, name)
}

func PostgresRepo(pkgPath string, module string) string {
	name := strcase.ToCamel(module)
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
	_, err := r.db.Exec("INSERT INTO %ss (title, body) VALUES ($1, $2)", %s.Title, %s.Body)
	return err
}

func (r *%sRepo) FindAll() ([]%s.%s, error) {
	rows, err := r.db.Query("SELECT id, title, body, created_at, updated_at, deleted_at FROM %ss")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []%s.%s
	for rows.Next() {
		var obj %s.%s
		if err := rows.Scan(&obj.ID, &obj.Title, &obj.Body, &obj.CreatedAt, &obj.UpdatedAt, &obj.DeletedAt); err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

func (r *%sRepo) FindByID(id string) (*%s.%s, error) {
	var obj %s.%s
	err := r.db.QueryRow("SELECT id, title, body, created_at, updated_at, deleted_at FROM %ss WHERE id=$1", id).Scan(&obj.ID, &obj.Title, &obj.Body, &obj.CreatedAt, &obj.UpdatedAt, &obj.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("not found")
		}
		return nil, err
	}
	return &obj, nil
}

func (r *%sRepo) Update(id string, %s *%s.%s) error {
	_, err := r.db.Exec("UPDATE %ss SET title=$1, body=$2, updated_at=NOW() WHERE id=$3", %s.Title, %s.Body, id)
	return err
}

func (r *%sRepo) Delete(id string) error {
	// Soft delete
	_, err := r.db.Exec("UPDATE %ss SET deleted_at=NOW() WHERE id=$1", id)
	// Hard delete
	// Delete the user from the database
	// This will remove the user permanently
	// and not just mark it as deleted.
	// _, err := r.db.Exec("DELETE FROM %ss WHERE id=$1", id)	
	return err
}`, pkgPath, module, module, module, module,
		module, module, module, name, module,
		module, module, module, module, name,
		module, module, name, module, name,
		module, module, name, module, name,
		module, module, module, module, name,
		module, module, module, module, module,
		module)
}
