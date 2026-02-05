package provider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	rdsTypes "github.com/aws/aws-sdk-go-v2/service/rds/types"
)

func FetchRDSInstances(ctx context.Context, cfg aws.Config) ([]rdsTypes.DBInstance, error) {
	rdsClient := rds.NewFromConfig(cfg)

	// Hardcoded por enquanto, ok para MVP
	const maxInstances = 20

	output, err := rdsClient.DescribeDBInstances(ctx, &rds.DescribeDBInstancesInput{
		MaxRecords: aws.Int32(maxInstances),
	})

	if err != nil {
		return nil, err
	}

	return output.DBInstances, nil
}
