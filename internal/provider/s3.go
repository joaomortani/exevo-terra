package provider

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// FetchBuckets lista todos os buckets da conta
func FetchBuckets(ctx context.Context, cfg aws.Config) ([]interface{}, error) {
	client := s3.NewFromConfig(cfg)

	fmt.Println("🔍 Listando S3 Buckets...")
	output, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, fmt.Errorf("erro ao listar buckets: %w", err)
	}

	// Conversão para []interface{} (necessário para o adapter genérico)
	result := make([]interface{}, len(output.Buckets))
	for i, b := range output.Buckets {
		result[i] = b
	}

	return result, nil
}
