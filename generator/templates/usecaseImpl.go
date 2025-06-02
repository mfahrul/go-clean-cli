package templates

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

func UsecaseImpl(pkgPath string, module string, fields []Field) string {
	name := strcase.ToCamel(module)

	createAssignments := ""
	for _, f := range fields {
		createAssignments += fmt.Sprintf("%s: dto.%s,\n", f.Name, f.Name)
	}

	responseAssignments := ""
	for _, f := range fields {
		responseAssignments += fmt.Sprintf("%s: obj.%s,\n", f.Name, f.Name)
	}

	updateLogic := ""
	for _, f := range fields {
		updateLogic += fmt.Sprintf(`if obj.%s != dto.%s && dto.%s != "" {
if !changed {
changed = true
}
obj.%s = dto.%s
}
`, f.Name, f.Name, f.Name, f.Name, f.Name)
	}

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
%s}
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
%s}
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
%s}
return result, nil
}

func (uc *%sUsecase) Update(id string, dto %s.UpdateDTO) (*%s.ResponseDTO, error) {
obj, err := uc.repo.FindByID(id)
if err != nil {
return nil, err
}
var changed = false
if obj == nil {
return nil, fmt.Errorf("object with id %%s not found", id)
}
%s
if !changed {
return nil, fmt.Errorf("no changes detected")
}
err = uc.repo.Update(id, obj)
if err != nil {
return nil, err
}
result := &%s.ResponseDTO{
%s}
return result, nil
}

func (uc *%sUsecase) Delete(id string) error {
return uc.repo.Delete(id)
}`, pkgPath, module, module, module, module,
		module, module, module, module, module,
		name, module, module, createAssignments,
		module, module, module, module, module,
		module, module, responseAssignments,
		module, module, module, responseAssignments,
		module, module, module, updateLogic,
		module, responseAssignments,
		module)
}
