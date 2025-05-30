package templates

import "fmt"

func UnitTest(pkgPath string, module string) string {
	return fmt.Sprintf(`package test

import (
	"testing"
//	"%s/internal/%s"
)

func TestCreate(t *testing.T) {
	// TODO: implement unit test for Create
}`, pkgPath, module)
}

func IntegrationTest(pkgPath string, module string) string {
	return fmt.Sprintf(`package test

import (
	"testing"
//	"database/sql"
	_ "github.com/lib/pq"
//	"%s/internal/%s"
)

func TestIntegrationCreate(t *testing.T) {
	// TODO: implement integration test for Create with real DB
}`, pkgPath, module)
}
