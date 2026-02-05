package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
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
	RunE:  runGenerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("filter", "f", "", "Filtra recursos pelo nome")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	resourceType, _ := cmd.Flags().GetString("resource")
	filter, _ := cmd.Flags().GetString("filter")

	config, globalConfig, err := loadConfig(resourceType)

	if err != nil {
		return err
	}

	fmt.Printf("☁️  Buscando recursos do tipo '%s' na AWS...\n", resourceType)
	rawResources, err := fetchResources(cmd.Context(), resourceType)
	if err != nil {
		return err
	}
	resourceDataList, err := adapter.BatchToMap(rawResources)
	if err != nil {
		return fmt.Errorf("erro ao adaptar recursos: %w", err)
	}

	filteredList := applyFilter(resourceDataList, config.PrimaryKey, filter)
	if len(filteredList) == 0 {
		return fmt.Errorf("nenhum recurso encontrado com o filtro '%s'", filter)
	}

	return writeOutput(filteredList, config, globalConfig, resourceType)

}

func loadConfig(resType string) (configuration.Resource, configuration.GlobalConfig, error) {
	fmt.Println("📖 Lendo exevo.yaml...")
	fullConfig, err := configuration.Load("exevo.yaml")
	if err != nil {
		return configuration.Resource{}, configuration.GlobalConfig{}, fmt.Errorf("falha ao ler config: %w", err)
	}
	resourceConfig, ok := fullConfig.Resources[resType]
	if !ok {
		return configuration.Resource{}, configuration.GlobalConfig{}, fmt.Errorf("recurso '%s' não encontrado", resType)
	}

	return resourceConfig, fullConfig.Global, nil
}

func fetchResources(ctx context.Context, resType string) ([]interface{}, error) {
	switch resType {
	case "rds":
		instances, err := provider.FetchRDSInstances(ctx, sharedAwsCfg)
		if err != nil {
			return nil, err
		}

		list := make([]interface{}, len(instances))
		for i, v := range instances {
			list[i] = v
		}
		return list, nil

	case "s3":
		return provider.FetchBuckets(ctx, sharedAwsCfg)

	default:
		return nil, fmt.Errorf("provider '%s' ainda não implementado via código", resType)
	}
}

func applyFilter(list []adapter.ResourceData, key, filter string) []adapter.ResourceData {
	if filter == "" {
		return list
	}

	var filtered []adapter.ResourceData
	for _, item := range list {
		// Type assertion segura
		if name, ok := item[key].(string); ok {
			if strings.Contains(name, filter) {
				filtered = append(filtered, item)
			}
		}
	}
	return filtered
}

func writeOutput(list []adapter.ResourceData, config configuration.Resource, globalConfig configuration.GlobalConfig, resType string) error {
	fmt.Printf("🚀 Gerando Terraform para %d recursos...\n", len(list))

	outputDir := filepath.Join("infra", resType)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("falha ao criar diretório output: %w", err)
	}

	mainFile := filepath.Join(outputDir, "main.tf")
	importsFile := filepath.Join(outputDir, "imports.tf")
	versionsFile := filepath.Join(outputDir, "versions.tf")

	// Gera Definição
	if err := generator.GenerateGeneric(list, config, mainFile); err != nil {
		return err
	}

	// Gera Imports (Igual)
	if err := generator.GenerateGenericImport(list, config, importsFile); err != nil {
		return err
	}

	// Gera Versions (NOVO!) 🆕
	fmt.Println("🏗️  Gerando Configuração Global (versions.tf)...")
	if err := generator.GenerateVersions(globalConfig, resType, versionsFile); err != nil {
		return fmt.Errorf("erro ao gerar versions.tf: %w", err)
	}

	fmt.Printf("✅ Sucesso! Arquivos gerados:\n - %s\n - %s\n - %s\n", mainFile, importsFile, versionsFile)
	return nil
}
