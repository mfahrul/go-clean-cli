package generator

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"crudgen/generator/templates"
)

func ReadModulePath() (string, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	return "", errors.New("module path not found in go.mod")
}

func Generate(module string, pkgPath string, fields []templates.Field) {
	fmt.Printf("🚀 Generating module: %s\n", module)
	appendRouteRegistration(module, pkgPath)
	// Create directories
	dirs := []string{
		fmt.Sprintf("internal/%s/delivery/http", module),
		fmt.Sprintf("internal/%s/repository/postgres", module),
		fmt.Sprintf("internal/%s/usecase", module),
		fmt.Sprintf("internal/%s/test", module),
	}
	for _, dir := range dirs {
		_ = os.MkdirAll(dir, os.ModePerm)
	}

	files := map[string]string{
		fmt.Sprintf("internal/%s/entity.go", module):                              templates.Entity(module, fields),
		fmt.Sprintf("internal/%s/dto.go", module):                                 templates.DTO(module, fields),
		fmt.Sprintf("internal/%s/repository.go", module):                          templates.Repository(module, fields),
		fmt.Sprintf("internal/%s/usecase.go", module):                             templates.Usecase(module, fields),
		fmt.Sprintf("internal/%s/usecase/%s_usecase.go", module, module):          templates.UsecaseImpl(pkgPath, module, fields),
		fmt.Sprintf("internal/%s/repository/postgres/%s_repo.go", module, module): templates.PostgresRepo(pkgPath, module, fields),
		fmt.Sprintf("internal/%s/delivery/http/handler.go", module):               templates.Handler(pkgPath, module),
		fmt.Sprintf("internal/%s/delivery/http/routes.go", module):                templates.Routes(pkgPath, module),
		fmt.Sprintf("internal/%s/test/unit_test.go", module):                      templates.UnitTest(pkgPath, module),
		fmt.Sprintf("internal/%s/test/integration_test.go", module):               templates.IntegrationTest(pkgPath, module),
	}

	for path, content := range files {
		_ = os.WriteFile(path, []byte(content), 0644)
		fmt.Println("📦 Created:", path)
	}

	fmt.Printf("✅ Module '%s' generated!\n", module)
}

func appendRouteRegistration(module string, pkgPath string) {
	fmt.Printf("🔗 Registering route for module: %s\n", module)

	routesFile := "core/routers/routers.go"
	line := fmt.Sprintf("_ \"%s/internal/%s/delivery/http\"", pkgPath, module)

	data, err := os.ReadFile(routesFile)
	if err != nil {
		fmt.Printf("❌ Skipped auto-registering route: %s not found\n", routesFile)
		os.Exit(1)
		return
	}
	content := string(data)
	if strings.Contains(content, line) {
		fmt.Printf("❌ Failed to auto-registering route: %s already registered\n", line)
		os.Exit(1)
		return // already registered
	}
	updated := strings.Replace(content, "// REGISTER ROUTES HERE", line+"\n\t// REGISTER ROUTES HERE", 1)
	_ = os.WriteFile(routesFile, []byte(updated), 0644)
	fmt.Println("🔗 Route registered in:", routesFile)
}
