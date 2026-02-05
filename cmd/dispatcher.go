package cmd

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/joaomortani/exevo-terra/internal/provider"
)

func FetchGenericResources(ctx context.Context, cfg aws.Config, resourceType string) ([]interface{}, error) {
	switch resourceType {
	case "rds":
		list, err := provider.FetchRDSInstances(ctx, sharedAwsCfg)
		if err != nil {
			return nil, err
		}
		// O genérico funciona automaticamente aqui
		return toInterfaceSlice(list), nil

	case "s3":
		return provider.FetchBuckets(ctx, cfg)
	case "ecs":
		list, err := provider.FetchECSServices(ctx, sharedAwsCfg)
		if err != nil {
			return nil, err
		}
		return toInterfaceSlice(list), nil

	case "ec2":
		rawList, err := provider.FetchEC2Instances(ctx, cfg)
		if err != nil {
			return nil, err
		}
		return toInterfaceSlice(rawList), nil

	default:
		return nil, fmt.Errorf("recurso '%s' ainda não suportado pelo provider", resourceType)
	}
}

func toInterfaceSlice[T any](input []T) []interface{} {
	result := make([]interface{}, len(input))
	for i, v := range input {
		result[i] = v
	}
	return result
}
