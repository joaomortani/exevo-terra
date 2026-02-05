package provider

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// FetchEC2Instances busca instâncias e retorna o wrapper genérico EnhancedResource
func FetchEC2Instances(ctx context.Context, cfg aws.Config) ([]EnhancedResource[types.Instance], error) {
	client := ec2.NewFromConfig(cfg)

	// 1. Primeiro coletamos os dados BRUTOS da AWS
	var rawInstances []types.Instance

	paginator := ec2.NewDescribeInstancesPaginator(client, &ec2.DescribeInstancesInput{})

	for paginator.HasMorePages() {
		// CORREÇÃO: Faltava chamar o NextPage para pegar os dados
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("erro ao listar EC2: %w", err)
		}

		// CORREÇÃO: EC2 retorna 'Reservations', precisamos iterar nelas
		for _, reservation := range page.Reservations {
			rawInstances = append(rawInstances, reservation.Instances...)
		}
	}

	// 2. Agora aplicamos o Enrichment usando o helper genérico
	enrichedList := EnrichSlice(rawInstances, func(inst types.Instance) map[string]interface{} {
		// Aqui definimos os campos sintéticos que queremos "flattar"
		return map[string]interface{}{
			"SimpleSecurityGroupIds": extractSgIds(inst.SecurityGroups),
			// Exemplo extra: garante string vazia em vez de nil pointer
			"SimpleSubnetId": aws.ToString(inst.SubnetId),
		}
	})

	return enrichedList, nil
}

// --- Helper Local ---

// extractSgIds transforma []GroupIdentifier em []string
func extractSgIds(groups []types.GroupIdentifier) []string {
	ids := make([]string, 0, len(groups))
	for _, g := range groups {
		if g.GroupId != nil {
			ids = append(ids, *g.GroupId)
		}
	}
	return ids
}
