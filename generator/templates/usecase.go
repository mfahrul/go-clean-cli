package templates

import (
	"fmt"
)

func Usecase(module string, fields []Field) string {
	return fmt.Sprintf(`package %s

	type Usecase interface {
		Create(dto CreateDTO) error
		FindAll() ([]ResponseDTO, error)
		FindByID(id string) (*ResponseDTO, error)
		Update(id string, dto UpdateDTO) (*ResponseDTO, error)
		Delete(id string) error
	}`, module)
}
