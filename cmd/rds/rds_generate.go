package rds

import (
	"github.com/joaomortani/exevo-terra/internal/adapter"
	"github.com/joaomortani/exevo-terra/internal/generator"
	"github.com/joaomortani/exevo-terra/internal/provider"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Gera blocos de código Terraform",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Reúso total: Usa a mesma config, mesma função de fetch
		instances, err := provider.FetchRDSInstances(cmd.Context(), sharedAwsCfg)
		if err != nil {
			return err
		}

		cleanList := adapter.BatchAwsToDomain(instances)
		outputPath := "./terraform.tf"
		return generator.GenerateFile(cleanList, outputPath)
	},
}

func init() {
	RdsCmd.AddCommand(generateCmd)
}
