package helpers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func LoadConfig(ctx context.Context, region string, profileName string) (aws.Config, error) {
	opts := []func(*config.LoadOptions) error{
		config.WithRegion(region),
	}

	if profileName != "" {
		opts = append(opts, config.WithSharedConfigProfile(profileName))
	}

	// Retorna o erro, não mata o processo
	return config.LoadDefaultConfig(ctx, opts...)
}
