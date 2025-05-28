package cmd

import (
	"fmt"
	"os"

	"crudgen/generator"

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
		generator.Generate(moduleName, projectModulePath)
	},
}

func Execute() {
	rootCmd.Flags().StringVarP(&moduleName, "module", "m", "", "Module name")
	_ = rootCmd.Execute()
}
