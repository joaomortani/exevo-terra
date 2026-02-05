package generator

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/joaomortani/exevo-terra/internal/adapter"
	"github.com/joaomortani/exevo-terra/internal/configuration"
	"github.com/zclconf/go-cty/cty"
)

// GenerateGeneric cria o HCL baseado APENAS no YAML e no Mapa de Dados
// Não sabe o que é RDS, S3 ou EC2. Sabe apenas seguir regras.
func GenerateGeneric(resources []adapter.ResourceData, config configuration.Resource, filename string) error {
	f := hclwrite.NewEmptyFile()
	rootBody := f.Body()

	for _, res := range resources {
		// 1. Descobre o nome do módulo (ex: identifier -> "meu-banco")
		// O YAML diz qual campo é a chave primária (ex: "DBInstanceIdentifier")
		pkField := config.PrimaryKey

		// Tenta pegar o valor do mapa genérico
		pkValue, ok := res[pkField].(string)
		if !ok || pkValue == "" {
			// Fallback de segurança se o campo não existir ou não for string
			pkValue = "unnamed-resource"
		}

		// Cria o bloco: module "NOME" { ... }
		moduleBlock := rootBody.AppendNewBlock("module", []string{pkValue})
		moduleBody := moduleBlock.Body()

		// 2. Define o Source do Módulo (Vem do YAML)
		moduleBody.SetAttributeValue("source", cty.StringVal(config.Source))

		// 3. Itera sobre os Mapeamentos do YAML (A Mágica acontece aqui ✨)
		// key = Nome da variável no Terraform (ex: instance_class)
		// field = Nome do campo na AWS (ex: DBInstanceClass)

		// Dica: Ordenamos as chaves para garantir que o arquivo gerado seja determinístico
		// (senão a cada execução a ordem das linhas muda)
		keys := make([]string, 0, len(config.Mappings))
		for k := range config.Mappings {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, key := range keys {
			awsField := config.Mappings[key]

			// Se o valor existir no mapa da AWS, escreve no Terraform
			if val, exists := res[awsField]; exists && val != nil {
				// Precisamos converter o interface{} do Go para cty.Value do HCL
				ctyVal := toCtyValue(val)
				moduleBody.SetAttributeValue(key, ctyVal)
			}
		}

		// 4. Adiciona valores estáticos (ex: terraform_managed = true)
		// Também ordenamos para manter o arquivo bonito
		staticKeys := make([]string, 0, len(config.Static))
		for k := range config.Static {
			staticKeys = append(staticKeys, k)
		}
		sort.Strings(staticKeys)

		for _, key := range staticKeys {
			val := config.Static[key]
			moduleBody.SetAttributeValue(key, toCtyValue(val))
		}

		// Adiciona uma linha em branco entre módulos
		rootBody.AppendNewline()
	}

	// 5. Salva no disco
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo: %w", err)
	}
	defer file.Close()

	_, err = f.WriteTo(file)
	if err != nil {
		return fmt.Errorf("erro ao escrever HCL: %w", err)
	}

	return nil
}

func GenerateGenericImport(resources []adapter.ResourceData, config configuration.Resource, filename string) error {
	f := hclwrite.NewEmptyFile()
	rootBody := f.Body()

	// Separa o tipo e o nome (ex: "aws_db_instance.this" -> ["aws_db_instance", "this"])
	parts := strings.Split(config.ResourceAddress, ".")
	if len(parts) != 2 {
		return fmt.Errorf("resource_address inválido no YAML. Use formato 'tipo.nome' (ex: aws_db_instance.this)")
	}
	resType := parts[0]
	resName := parts[1]

	for _, res := range resources {
		// Pega o ID (Primary Key)
		pkValue, ok := res[config.PrimaryKey].(string)
		if !ok || pkValue == "" {
			continue // Pula recursos sem ID
		}

		importBlock := rootBody.AppendNewBlock("import", nil)
		importBody := importBlock.Body()

		// 1. ID na AWS
		importBody.SetAttributeValue("id", cty.StringVal(pkValue))

		// 2. Caminho no Terraform (Traversal)
		// Gerando: module.{pkValue}.{resType}.{resName}
		toTokens := hclwrite.Tokens{
			{Type: hclsyntax.TokenIdent, Bytes: []byte("module")},
			{Type: hclsyntax.TokenDot, Bytes: []byte(".")},
			{Type: hclsyntax.TokenIdent, Bytes: []byte(pkValue)},
			{Type: hclsyntax.TokenDot, Bytes: []byte(".")},
			{Type: hclsyntax.TokenIdent, Bytes: []byte(resType)},
			{Type: hclsyntax.TokenDot, Bytes: []byte(".")},
			{Type: hclsyntax.TokenIdent, Bytes: []byte(resName)},
		}

		importBody.SetAttributeRaw("to", toTokens)
		rootBody.AppendNewline()
	}

	// Salva arquivo
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = f.WriteTo(file)
	return err

}

// Helper para converter tipos primitivos do Go para CTY (Terraform)
// O json.Unmarshal converte números para float64 por padrão, precisamos tratar isso.
func toCtyValue(v interface{}) cty.Value {
	switch val := v.(type) {
	case string:
		return cty.StringVal(val)
	case int:
		return cty.NumberIntVal(int64(val))
	case int32:
		return cty.NumberIntVal(int64(val))
	case int64:
		return cty.NumberIntVal(val)
	case float64:
		// JSON numbers são float64. Geralmente em infra queremos Int (portas, GBs).
		// Se precisar de float real (ex: CPU credits), pode usar NumberFloatVal.
		// Por segurança, vamos converter para Int se não tiver decimal.
		return cty.NumberIntVal(int64(val))
	case bool:
		return cty.BoolVal(val)
	default:
		// Fallback para string se não souber o que é
		return cty.StringVal(fmt.Sprintf("%v", val))
	}
}
