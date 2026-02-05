package rds

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joaomortani/exevo-terra/internal/adapter"
	"github.com/joaomortani/exevo-terra/internal/provider"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Mostra a estrutura JSON crua de um recurso para ajudar no mapeamento YAML",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 1. Busca os recursos (Reúso do Provider)
		instances, err := provider.FetchRDSInstances(cmd.Context(), sharedAwsCfg)
		if err != nil {
			return err
		}

		if len(instances) == 0 {
			fmt.Println("Nenhum recurso encontrado para inspecionar.")
			return nil
		}

		sample := instances[0]

		dataMap, err := adapter.ToMap(sample)
		if err != nil {
			return err
		}

		// 4. Pretty Print do JSON
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")

		fmt.Println("// --- Estrutura disponível para mapeamento no exevo.yaml ---")
		fmt.Println("// Copie as chaves (Keys) abaixo para o seu arquivo de configuração.")

		return encoder.Encode(dataMap)
	},
}

func init() {
	RdsCmd.AddCommand(inspectCmd)
}
