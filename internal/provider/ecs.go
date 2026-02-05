package provider

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

// FetchECSServices agora retorna EnhancedResource
func FetchECSServices(ctx context.Context, cfg aws.Config) ([]EnhancedResource[types.Service], error) {
	client := ecs.NewFromConfig(cfg)
	var rawServices []types.Service

	// 1. Listar Clusters
	clustersPag := ecs.NewListClustersPaginator(client, &ecs.ListClustersInput{})

	for clustersPag.HasMorePages() {
		clusterPage, err := clustersPag.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("erro ao listar clusters ECS: %w", err)
		}

		for _, clusterArn := range clusterPage.ClusterArns {
			// 2. Listar Services dentro do Cluster
			servicesPag := ecs.NewListServicesPaginator(client, &ecs.ListServicesInput{
				Cluster: aws.String(clusterArn),
			})

			for servicesPag.HasMorePages() {
				servicePage, err := servicesPag.NextPage(ctx)
				if err != nil {
					// Log e continua ou falha. Vamos falhar safe.
					return nil, fmt.Errorf("erro ao listar services: %w", err)
				}

				if len(servicePage.ServiceArns) == 0 {
					continue
				}

				// 3. Describe Services (Batch de 10 é o limite da AWS)
				// Para MVP assumimos < 10 por página. Para Prod precisaria de chunking.
				descResp, err := client.DescribeServices(ctx, &ecs.DescribeServicesInput{
					Cluster:  aws.String(clusterArn),
					Services: servicePage.ServiceArns,
				})
				if err != nil {
					return nil, fmt.Errorf("erro ao descrever services: %w", err)
				}

				rawServices = append(rawServices, descResp.Services...)
			}
		}
	}

	// 4. Enrichment: Trazendo dados de rede para a superfície
	return EnrichSlice(rawServices, func(svc types.Service) map[string]interface{} {
		extras := make(map[string]interface{})

		// Campos padrão vazios
		extras["SimpleSubnetIds"] = []string{}
		extras["SimpleSecurityGroupIds"] = []string{}
		extras["IsFargate"] = false

		// Se tiver config de rede (Fargate ou awsvpc), extrai os dados
		if svc.NetworkConfiguration != nil && svc.NetworkConfiguration.AwsvpcConfiguration != nil {
			extras["SimpleSubnetIds"] = svc.NetworkConfiguration.AwsvpcConfiguration.Subnets
			extras["SimpleSecurityGroupIds"] = svc.NetworkConfiguration.AwsvpcConfiguration.SecurityGroups

			if svc.LaunchType == types.LaunchTypeFargate {
				extras["IsFargate"] = true
			}
		}

		return extras
	}), nil
}
