package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Cria um arquivo exevo.yaml de exemplo no diretório atual",
	RunE: func(cmd *cobra.Command, args []string) error {
		force, _ := cmd.Flags().GetBool("force")
		filename := "exevo.yaml"

		if _, err := os.Stat(filename); err == nil {
			if !force {
				return fmt.Errorf("o arquivo '%s' já existe. Use --force para sobrescrever", filename)
			}
			fmt.Println("⚠️  Sobrescrevendo arquivo exevo.yaml existente...")
		}

		fmt.Println("🌱 Criando exevo.yaml de exemplo...")
		if err := os.WriteFile(filename, []byte(yamlTemplate), 0644); err != nil {
			return fmt.Errorf("erro ao escrever arquivo: %w", err)
		}

		fmt.Println("✅ Sucesso! Agora edite o 'exevo.yaml' e rode:")
		fmt.Println("   exevo-terra generate --resource s3")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolP("force", "f", false, "Sobrescreve o arquivo se já existir")
}

// O Template "Baterias Inclusas"
// Mantemos aqui como const para o binário ser self-contained (sem depender de arquivos externos)
const yamlTemplate = `version: "1"

# ------------------------------------------------------------------
# 🌍 GLOBAL CONFIGURATION
# Configurações do Terraform que serão aplicadas no versions.tf
# ------------------------------------------------------------------
global:
  terraform_version: ">= 1.5.0"
  
  backend:
    type: "s3"
    config:
      bucket: "meu-bucket-terraform-state"
      # {{RESOURCE}} é substituído dinamicamente pelo nome do recurso (ex: rds, s3)
      key: "exevo-terra/{{RESOURCE}}/terraform.tfstate"
      region: "us-east-1"
      encrypt: true

  providers:
    aws:
      source: "hashicorp/aws"
      version: ">= 5.0"

# ------------------------------------------------------------------
# 📦 RESOURCES CONFIGURATION
# Mapeamento de recursos da AWS para Módulos Terraform
# ------------------------------------------------------------------
resources:

  # Exemplo: Amazon RDS
  rds:
    # Módulo Terraform que será utilizado (pode ser local ou do registry)
    source: "./modules/rds-padrao"
    
    # Campo da API AWS usado como identificador único (Nome do Módulo)
    primary_key: "DBInstanceIdentifier"
    
    # Endereço do recurso dentro do módulo (para importação)
    resource_address: "aws_db_instance.this" 

    # Mapeamento: Variável Terraform <= Campo AWS (Case Sensitive)
    # Use 'exevo-terra inspect' para ver os campos disponíveis
    mappings:
      identifier: "DBInstanceIdentifier"
      instance_class: "DBInstanceClass"
      engine: "Engine"
      allocated_storage: "AllocatedStorage"
    
    # Valores estáticos (forçados no código gerado)
    static:
      terraform_managed: true
      environment: "production"

  # Exemplo: Amazon S3
  s3:
    source: "terraform-aws-modules/s3-bucket/aws"
    primary_key: "Name"
    resource_address: "aws_s3_bucket.this"
    
    mappings:
      bucket: "Name"
    
    static:
      acl: "private"
      control_object_ownership: true
      object_ownership: "ObjectWriter"
`
