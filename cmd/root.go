package cmd

import (
	"fmt"
	"os"
	"strings"

	"crudgen/generator"
	"crudgen/generator/templates"

	"github.com/spf13/cobra"
)

var moduleName string

var rootCmd = &cobra.Command{
	Use:   "crudgen",
	Short: "CRUD Generator CLI",
	Run: func(cmd *cobra.Command, args []string) {
		if moduleName == "" {
			fmt.Println("❌ Module name is required. Use -m flag.")
			os.Exit(1)
		}
		projectModulePath, err := generator.ReadModulePath()
		if err != nil {
			fmt.Println("❌ Failed to read go.mod:", err)
			os.Exit(1)
		}
		fields := []templates.Field{}
		reader := os.Stdin
		fmt.Println("Enter struct fields (name type), one per line. Leave empty to finish:")
		for {
			var name, typ string
			fmt.Print("Field name: ")
			fmt.Fscanln(reader, &name)
			if name == "" {
				break
			}
			fmt.Print("Field type: ")
			fmt.Fscanln(reader, &typ)
			if typ == "" {
				break
			}
			fields = append(fields, templates.Field{Name: name, Type: typ})
		}
		generator.Generate(moduleName, projectModulePath, fields)
	},
}

// ToSnakeCase converts a string to snake_case format
func ToSnakeCase(str string) string {
	var result string
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result += "_"
		}
		result += string(r)
	}
	return strings.ToLower(result)
}

func Execute() {
	rootCmd.Flags().StringVarP(&moduleName, "module", "m", "", "Module name")
	_ = rootCmd.Execute()
}
