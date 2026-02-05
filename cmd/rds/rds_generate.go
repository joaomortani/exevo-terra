package rds

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
		// 1. Carrega a Configuração YAML
		fmt.Println("📖 Lendo exevo.yaml...")
		config, err := configuration.Load("exevo.yaml")
		if err != nil {
			return fmt.Errorf("falha ao ler config: %w", err)
		}

		// Valida se tem configuração para RDS
		rdsConfig, ok := config.Resources["rds"]
		if !ok {
			return fmt.Errorf("nenhuma configuração 'rds' encontrada no exevo.yaml")
		}

		// 2. Busca dados na AWS (Usando a credencial compartilhada do pai)
		fmt.Println("☁️  Buscando recursos na AWS...")
		instances, err := provider.FetchRDSInstances(cmd.Context(), sharedAwsCfg)
		if err != nil {
			return err
		}

		// 3. O Pulo do Gato: Converter []Concrete para []Interface
		// Necessário para o adapter genérico funcionar
		genericList := make([]interface{}, len(instances))
		for i, v := range instances {
			genericList[i] = v
		}

		// 4. Converte para Mapa Genérico (JSON Hack)
		resourceDataList, err := adapter.BatchToMap(genericList)
		if err != nil {
			return err
		}

		// 5. Filtro Dinâmico (Opcional)
		// Usa a Primary Key definida no YAML para filtrar
		filter, _ := cmd.Flags().GetString("filter")
		var filteredList []adapter.ResourceData

		if filter != "" {
			for _, res := range resourceDataList {
				// Tenta pegar o valor da chave primária (ex: DBInstanceIdentifier)
				if name, ok := res[rdsConfig.PrimaryKey].(string); ok {
					if strings.Contains(name, filter) {
						filteredList = append(filteredList, res)
					}
				}
			}
		} else {
			filteredList = resourceDataList
		}

		// 6. Gera o Código HCL usando o Motor Genérico
		fmt.Printf("🚀 Gerando Terraform para %d recursos...\n", len(filteredList))

		outputFile := "rds_dynamic.tf"
		if err := generator.GenerateGeneric(filteredList, rdsConfig, outputFile); err != nil {
			return err
		}

		fmt.Println("🔗 Gerando Imports Dinâmicos...")
		if err := generator.GenerateGenericImport(filteredList, rdsConfig, "imports.tf"); err != nil {
			return err
		}

		fmt.Printf("✅ Sucesso! Arquivo '%s' gerado.\n", outputFile)
		return nil
	},
}

func init() {
	RdsCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("filter", "f", "", "Filtra recursos pelo nome")
}
