package rds

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/joaomortani/exevo-terra/internal/helpers"
	"github.com/spf13/cobra"
)

var sharedAwsCfg aws.Config

var RdsCmd = &cobra.Command{
	Use:   "rds",
	Short: "Gerencia recursos RDS",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		region, _ := cmd.Flags().GetString("region")
		profile, _ := cmd.Flags().GetString("profile")

		ctx := context.Background()

		cfg, err := helpers.LoadConfig(ctx, region, profile)
		if err != nil {
			log.Fatalf("Erro ao carregar AWS: %v", err)
		}
		sharedAwsCfg = cfg

		return nil

	},
}
