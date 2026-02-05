package provider

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

func FetchRDSInstances(ctx context.Context, cfg aws.Config) ([]EnhancedResource[types.DBInstance], error) {
	client := rds.NewFromConfig(cfg)
	var rawInstances []types.DBInstance

	paginator := rds.NewDescribeDBInstancesPaginator(client, &rds.DescribeDBInstancesInput{})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("erro ao listar RDS: %w", err)
		}
		rawInstances = append(rawInstances, page.DBInstances...)
	}

	// Enrichment: Simplificando listas de objetos
	return EnrichSlice(rawInstances, func(db types.DBInstance) map[string]interface{} {
		return map[string]interface{}{
			// Transforma [{VpcSecurityGroupId: "sg-1", ...}] em ["sg-1"]
			"SimpleSecurityGroupIds": extractRdsSgIds(db.VpcSecurityGroups),

			// Transforma DBSubnetGroup em lista de IDs de subnet
			"SimpleSubnetIds": extractRdsSubnetIds(db.DBSubnetGroup),
		}
	}), nil
}

// --- Helpers Locais ---

func extractRdsSgIds(groups []types.VpcSecurityGroupMembership) []string {
	ids := make([]string, 0, len(groups))
	for _, g := range groups {
		if g.VpcSecurityGroupId != nil {
			ids = append(ids, *g.VpcSecurityGroupId)
		}
	}
	return ids
}

func extractRdsSubnetIds(group *types.DBSubnetGroup) []string {
	if group == nil {
		return []string{}
	}
	ids := make([]string, 0, len(group.Subnets))
	for _, s := range group.Subnets {
		if s.SubnetIdentifier != nil {
			ids = append(ids, *s.SubnetIdentifier)
		}
	}
	return ids
}
