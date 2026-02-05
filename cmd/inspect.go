package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joaomortani/exevo-terra/internal/adapter"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Mostra a estrutura JSON crua de um recurso para ajudar no mapeamento YAML",
	RunE: func(cmd *cobra.Command, args []string) error {
		resourceType, _ := cmd.Flags().GetString("resource")

		resource, err := FetchGenericResources(cmd.Context(), sharedAwsCfg, resourceType)
		if err != nil {
			return err
		}

		if len(resource) == 0 {
			fmt.Println("Nenhum recurso encontrado para inspecionar.")
			return nil
		}
		sample := resource[0]
		dataMap, err := adapter.ToMap(sample)
		if err != nil {
			return err
		}
		fmt.Println("// --- Estrutura disponível para mapeamento no exevo.yaml ---")
		fmt.Printf("// Recurso: %s\n", resourceType)
		fmt.Println("// Copie as chaves (Keys) abaixo para o seu 'mappings':")
		fmt.Println("")

		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")

		return encoder.Encode(dataMap)

	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}
