package templates

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

func Usecase(module string) string {
	return fmt.Sprintf(`package %s

type Usecase interface {
	Create(dto CreateDTO) (error)
	FindAll() ([]ResponseDTO, error)
	FindByID(id string) (*ResponseDTO, error)
	Update(id string, dto UpdateDTO) (*ResponseDTO, error)
	Delete(id string) error
}`, module)
}

func UsecaseImpl(pkgPath string, module string) string {
	name := strcase.ToCamel(module)
	return fmt.Sprintf(`package usecase

import (
	"fmt"
	"%s/internal/%s"
)
type %sUsecase struct {
	repo %s.Repository
}

func New(repo %s.Repository) %s.Usecase {
	return &%sUsecase{repo: repo}
}

func (uc *%sUsecase) Create(dto %s.CreateDTO) (error) {
	obj := &%s.%s{
		Title: dto.Title,
		Body:  dto.Body,
	}
	return uc.repo.Create(obj)
}

func (uc *%sUsecase) FindAll() ([]%s.ResponseDTO, error) {
	%ss, err := uc.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var responseDTOs []%s.ResponseDTO
	for _, obj := range %ss {
		respDTO := &%s.ResponseDTO{
			ID:    obj.ID,
			Title: obj.Title,
			Body:  obj.Body,
			CreatedAt: obj.CreatedAt,
			UpdatedAt: obj.UpdatedAt,
			DeletedAt: obj.DeletedAt,
		}
		responseDTOs = append(responseDTOs, *respDTO)
	}
	return responseDTOs, nil
}

func (uc *%sUsecase) FindByID(id string) (*%s.ResponseDTO, error) {
	obj, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	result := &%s.ResponseDTO{
		ID:    obj.ID,
		Title: obj.Title,
		Body:  obj.Body,
		CreatedAt: obj.CreatedAt,
		UpdatedAt: obj.UpdatedAt,
		DeletedAt: obj.DeletedAt,
	}
	return result, nil
}

func (uc *%sUsecase) Update(id string, dto %s.UpdateDTO) (*%s.ResponseDTO, error) {
	obj, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	var changed = false
	if obj == nil {
		return nil, fmt.Errorf("object with id $s not found", id)
	}
	if obj.Title != dto.Title && dto.Title != "" {
		if !changed {
			changed = true
		}
		obj.Title = dto.Title
	}
	if obj.Body != dto.Body && dto.Body != "" {
		if !changed {
			changed = true
		}
		obj.Body = dto.Body
	}
	if !changed {
		return nil, fmt.Errorf("no changes detected")
	}
	err = uc.repo.Update(id, obj)
	if err != nil {
		return nil, err
	}
	result := &%s.ResponseDTO{
		ID:    obj.ID,
		Title: obj.Title,
		Body:  obj.Body,
		CreatedAt: obj.CreatedAt,
		UpdatedAt: obj.UpdatedAt,
		DeletedAt: obj.DeletedAt,
	}
	return result, nil
}

func (uc *%sUsecase) Delete(id string) error {
	return uc.repo.Delete(id)
}`, pkgPath, module, module, module, module,
		module, module, module, module, module,
		name, module, module, module, module,
		module, module, module, module, module,
		module, module, module, module, module)
}
