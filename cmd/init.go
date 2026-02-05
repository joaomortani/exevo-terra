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
      # {{RESOURCE}} é substituído dinamicamente (ex: rds, s3, ec2)
      key: "exevo-terra/{{RESOURCE}}/terraform.tfstate"
      region: "us-east-1"
      encrypt: true
      dynamodb_table: "terraform-lock"

  providers:
    aws:
      source: "hashicorp/aws"
      version: ">= 5.0"

# ------------------------------------------------------------------
# 📦 RESOURCES CONFIGURATION
# ------------------------------------------------------------------
resources:

  # --- 🗄️ DATABASE (RDS) ---
  rds:
    source: "./modules/rds" # Aponte para seu módulo
    primary_key: "DBInstanceIdentifier"
    resource_address: "aws_db_instance.this" 
    
    mappings:
      identifier: "DBInstanceIdentifier"
      instance_class: "DBInstanceClass"
      engine: "Engine"
      allocated_storage: "AllocatedStorage"
      
      # ✨ Campos Enriquecidos (Exevo Terra Magic)
      # O Exevo Terra achata listas complexas para facilitar sua vida:
      vpc_security_group_ids: "SimpleSecurityGroupIds"
      subnet_ids: "SimpleSubnetIds"
    
    static:
      terraform_managed: true

  # --- 💻 COMPUTE (EC2) ---
  ec2:
    source: "terraform-aws-modules/ec2-instance/aws"
    primary_key: "InstanceId"
    resource_address: "aws_instance.this"
    
    mappings:
      name: "Tags.Name" # Busca valor da tag Name automaticamente
      instance_type: "InstanceType"
      ami: "ImageId"
      
      # ✨ Campos Enriquecidos
      vpc_security_group_ids: "SimpleSecurityGroupIds"
      subnet_id: "SimpleSubnetId"
    
    static:
      monitoring: true

  # --- 🐳 CONTAINERS (ECS Service) ---
  ecs:
    source: "terraform-aws-modules/ecs/aws//modules/service"
    primary_key: "ServiceName"
    resource_address: "aws_ecs_service.this"
    
    mappings:
      name: "ServiceName"
      cluster_arn: "ClusterArn"
      desired_count: "DesiredCount"
      
      # ✨ Campos Enriquecidos (Busca de Fargate/AwsvpcConfig)
      subnet_ids: "SimpleSubnetIds"
      security_group_ids: "SimpleSecurityGroupIds"
    
    static:
      ignore_task_definition_changes: true

  # --- ☁️ STORAGE (S3) ---
  s3:
    source: "terraform-aws-modules/s3-bucket/aws"
    primary_key: "Name"
    resource_address: "aws_s3_bucket.this"
    mappings:
      bucket: "Name"
    static:
      acl: "private"
`
