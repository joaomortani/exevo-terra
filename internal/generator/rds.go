package generator

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/joaomortani/exevo-terra/internal/adapter" // Importe seu adapter
	"github.com/zclconf/go-cty/cty"
)

func buildHCL(instances []adapter.RdsInstance) *hclwrite.File {
	f := hclwrite.NewEmptyFile()
	rootBody := f.Body()

	fmt.Printf("// Gerando configuração para %d recursos...\n\n", len(instances))
	for _, inst := range instances {
		moduleBlock := rootBody.AppendNewBlock("module", []string{inst.Name})
		moduleBody := moduleBlock.Body()
		// 3. Define os atributos básicos
		// Ajuste o "source" para o caminho real do módulo da sua empresa
		moduleBody.SetAttributeValue("source", cty.StringVal("./modules/rds-padrao"))

		moduleBody.SetAttributeValue("identifier", cty.StringVal(inst.Name))
		moduleBody.SetAttributeValue("engine", cty.StringVal(inst.Engine))
		moduleBody.SetAttributeValue("engine_version", cty.StringVal(inst.EngineVersion))
		moduleBody.SetAttributeValue("instance_class", cty.StringVal(inst.InstanceClass))
		moduleBody.SetAttributeValue("allocated_storage", cty.NumberIntVal(int64(inst.StorageGB)))
		moduleBody.SetAttributeValue("security_group_ids", cty.ListVal([]cty.Value{cty.StringVal(inst.InstanceSG)}))

		// Exemplo de condicional (Booleano)
		moduleBody.SetAttributeValue("multi_az", cty.BoolVal(inst.MultiAZ))

		// 4. Tratamento de Atributos Especiais (Storage Type)
		if inst.StorageType != "" {
			moduleBody.SetAttributeValue("storage_type", cty.StringVal(inst.StorageType))
		}

		// Adiciona uma linha em branco entre módulos para ficar bonito
		rootBody.AppendNewline()
	}

	return f
}

func PlanHCL(instances []adapter.RdsInstance) error {
	f := buildHCL(instances)

	fmt.Println("// --- Terraform Plan (Preview) ---")
	_, err := f.WriteTo(os.Stdout)
	return err
}

func GenerateFile(instances []adapter.RdsInstance, filename string) error {
	f := buildHCL(instances)

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo: %w", err)
	}
	defer file.Close()

	// Escreve os bytes no arquivo
	_, err = f.WriteTo(file)
	if err != nil {
		return fmt.Errorf("erro ao escrever HCL: %w", err)
	}

	fmt.Printf("✅ Arquivo gerado com sucesso: %s\n", filename)
	return nil

}
