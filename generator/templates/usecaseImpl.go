package templates

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

func UsecaseImpl(pkgPath string, module string, fields []Field) string {
	name := strcase.ToCamel(module)

	createAssignments := ""
	responseAssignments := ""
	updateLogic := ""

	for _, f := range fields {
		f.Name = strcase.ToCamel(f.Name)
		createAssignments += fmt.Sprintf("%s: dto.%s,\n\t", f.Name, f.Name)
		responseAssignments += fmt.Sprintf("%s: obj.%s,\n\t", f.Name, f.Name)
		updateLogic += fmt.Sprintf(`if dto.%s != nil && obj.%s != *dto.%s {
if !changed {
changed = true
}
obj.%s = *dto.%s
}
`, f.Name, f.Name, f.Name, f.Name, f.Name)
	}

	return fmt.Sprintf(`package usecase

import (
    "fmt"
    "math"
    "%s/internal/%s"
)

type %s_Usecase struct {
    repo %s.Repository
}

func New(repo %s.Repository) %s.Usecase {
    return &%s_Usecase{repo: repo}
}

func (uc *%s_Usecase) Create(dto %s.CreateDTO) error {
    obj := &%s.%s{
        %s
    }
    return uc.repo.Create(obj)
}

func (uc *%s_Usecase) FindAll(query *%s.PaginationQuery) (*%s.PaginationResponse, error) {
    // Set defaults if not provided
    if query.Page < 1 {
        query.Page = 1
    }
    if query.Limit < 1 {
        query.Limit = 10
    }
    if query.Order == "" {
        query.Order = "desc"
    }
    if query.Sort == "" {
        query.Sort = "created_at"
    }

    // Get paginated results
    %ss, total, err := uc.repo.FindAll(
        query.Page,
        query.Limit,
        query.Sort,
        query.Order,
        query.Search,
    )
    if err != nil {
        return nil, err
    }

    // Calculate pagination metadata
    totalPages := int(math.Ceil(float64(total) / float64(query.Limit)))
    hasNext := query.Page < totalPages
    hasPrev := query.Page > 1

    nextPage := query.Page + 1
    if !hasNext {
        nextPage = 0
    }

    prevPage := query.Page - 1
    if !hasPrev {
        prevPage = 0
    }

    // Convert entities to DTOs
    var responseDTOs []%s.ResponseDTO
    for _, obj := range %ss {
        respDTO := &%s.ResponseDTO{
            %s
        }
        responseDTOs = append(responseDTOs, *respDTO)
    }

    // Create pagination response
    response := &%s.PaginationResponse{
        Data: responseDTOs,
        Meta: %s.PaginationMeta{
            CurrentPage:  query.Page,
            TotalPages:   totalPages,
            TotalItems:   total,
            ItemsPerPage: query.Limit,
            HasNext:      hasNext,
            HasPrev:      hasPrev,
            NextPage:     nextPage,
            PrevPage:     prevPage,
        },
    }

    return response, nil
}

func (uc *%s_Usecase) FindByID(id string) (*%s.ResponseDTO, error) {
    obj, err := uc.repo.FindByID(id)
    if err != nil {
        return nil, err
    }
    result := &%s.ResponseDTO{
        %s
    }
    return result, nil
}

func (uc *%s_Usecase) Update(id string, dto %s.UpdateDTO) (*%s.ResponseDTO, error) {
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
        %s
    }
    return result, nil
}

func (uc *%s_Usecase) Delete(id string) error {
    return uc.repo.Delete(id)
}`, pkgPath, module, module, module, module,
		module, module, module, module, module,
		name, createAssignments,
		module, module, module,
		module, module, module, module, responseAssignments,
		module, module, module, module, module, responseAssignments,
		module, module, module, updateLogic,
		module, responseAssignments,
		module)
}
