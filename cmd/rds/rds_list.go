package rds

import (
	"fmt"

	"github.com/joaomortani/exevo-terra/internal/adapter"
	"github.com/joaomortani/exevo-terra/internal/provider"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista instâncias RDS encontradas",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 1. Usa a configuração que o PAI já carregou
		// Não precisa chamar LoadConfig de novo!
		instances, err := provider.FetchRDSInstances(cmd.Context(), sharedAwsCfg)
		if err != nil {
			return err
		}

		fmt.Printf("🔍 Encontradas %d instâncias:\n", len(instances))
		for _, instance := range instances {
			dto := adapter.AWSInstanceToConfig(instance)
			fmt.Printf("DB: %s | Storage: %d GB\n\n", dto.Name, dto.StorageGB)
		}

		return nil
	},
}

func init() {
	// Registra o filho no pai
	RdsCmd.AddCommand(listCmd)
}
