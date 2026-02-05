package generator

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/joaomortani/exevo-terra/internal/configuration"
	"github.com/zclconf/go-cty/cty"
)

// GenerateVersions cria o arquivo versions.tf com backend e providers
func GenerateVersions(config configuration.GlobalConfig, resourceName string, filename string) error {
	f := hclwrite.NewEmptyFile()
	rootBody := f.Body()

	// 1. Bloco terraform { ... }
	tfBlock := rootBody.AppendNewBlock("terraform", nil)
	tfBody := tfBlock.Body()

	// 1.1 Versão do Terraform
	if config.TerraformVersion != "" {
		tfBody.SetAttributeValue("required_version", cty.StringVal(config.TerraformVersion))
	}

	// 1.2 Bloco backend "tipo" { ... }
	if config.Backend.Type != "" {
		backendBlock := tfBody.AppendNewBlock("backend", []string{config.Backend.Type})
		backendBody := backendBlock.Body()

		// Ordenar chaves para determinismo
		keys := make([]string, 0, len(config.Backend.Config))
		for k := range config.Backend.Config {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			val := config.Backend.Config[k]

			// Lógica de Substituição Dinâmica (Magic Variable) 🎩
			// Se o valor for string e tiver {{RESOURCE}}, troca pelo nome atual (ex: rds)
			if strVal, ok := val.(string); ok {
				if strings.Contains(strVal, "{{RESOURCE}}") {
					val = strings.ReplaceAll(strVal, "{{RESOURCE}}", resourceName)
				}
			}

			backendBody.SetAttributeValue(k, toCtyValue(val))
		}
	}

	// 1.3 Bloco required_providers { ... }
	if len(config.Providers) > 0 {
		reqProvidersBlock := tfBody.AppendNewBlock("required_providers", nil)
		reqProvidersBody := reqProvidersBlock.Body()

		// Ordenar providers
		pKeys := make([]string, 0, len(config.Providers))
		for k := range config.Providers {
			pKeys = append(pKeys, k)
		}
		sort.Strings(pKeys)

		for _, pName := range pKeys {
			pConfig := config.Providers[pName]
			// Cria um objeto map dentro do HCL: aws = { source = "...", version = "..." }
			providerMap := map[string]cty.Value{
				"source":  cty.StringVal(pConfig.Source),
				"version": cty.StringVal(pConfig.Version),
			}
			reqProvidersBody.SetAttributeValue(pName, cty.ObjectVal(providerMap))
		}
	}

	// Salva no disco
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("erro ao criar versions.tf: %w", err)
	}
	defer file.Close()

	_, err = f.WriteTo(file)
	return err
}
