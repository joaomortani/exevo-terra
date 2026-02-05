package cmd

import (
	"fmt"
	"strings"

	"github.com/joaomortani/exevo-terra/internal/adapter"
	"github.com/joaomortani/exevo-terra/internal/configuration"
	"github.com/joaomortani/exevo-terra/internal/generator"
	"github.com/joaomortani/exevo-terra/internal/provider"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Gera código Terraform baseado no exevo.yaml",
	RunE: func(cmd *cobra.Command, args []string) error {

		resourceType, _ := cmd.Flags().GetString("resource")

		fmt.Println("📖 Lendo exevo.yaml...")
		config, err := configuration.Load("exevo.yaml")
		if err != nil {
			return fmt.Errorf("falha ao ler config: %w", err)
		}

		resourceConfig, ok := config.Resources[resourceType]
		if !ok {
			return fmt.Errorf("nenhuma configuração '%s' encontrada no exevo.yaml", resourceType)
		}

		var rawList []interface{}

		switch resourceType {
		case "rds":
			instances, err := provider.FetchRDSInstances(cmd.Context(), sharedAwsCfg)
			if err != nil {
				return err
			}

			rawList = make([]interface{}, len(instances))
			for i, v := range instances {
				rawList[i] = v
			}

		default:
			return fmt.Errorf("provider '%s' ainda não implementado no código Go", resourceType)
		}

		resourceDataList, err := adapter.BatchToMap(rawList)
		if err != nil {
			return err
		}

		// 5. Filtro Dinâmico
		filter, _ := cmd.Flags().GetString("filter")
		var filteredList []adapter.ResourceData

		if filter != "" {
			for _, res := range resourceDataList {
				// Usa a PrimaryKey definida no YAML para filtrar
				if name, ok := res[resourceConfig.PrimaryKey].(string); ok {
					if strings.Contains(name, filter) {
						filteredList = append(filteredList, res)
					}
				}
			}
		} else {
			filteredList = resourceDataList
		}

		if len(filteredList) == 0 {
			return fmt.Errorf("nenhum recurso encontrado com o filtro '%s'", filter)
		}

		// 6. Geração de Arquivos
		fmt.Printf("🚀 Gerando Terraform para %d recursos...\n", len(filteredList))

		// Nomes de arquivo dinâmicos
		outputFile := fmt.Sprintf("%s.tf", resourceType)
		importFile := fmt.Sprintf("import_%s.tf", resourceType)

		// Gera a Definição
		if err := generator.GenerateGeneric(filteredList, resourceConfig, outputFile); err != nil {
			return err
		}

		// Gera os Imports
		fmt.Println("🔗 Gerando Imports Dinâmicos...")
		if err := generator.GenerateGenericImport(filteredList, resourceConfig, importFile); err != nil {
			return err
		}

		fmt.Printf("✅ Sucesso! Arquivos gerados:\n - %s\n - %s\n", outputFile, importFile)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("filter", "f", "", "Filtra recursos pelo nome")
}
