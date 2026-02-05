/*
Copyright © 2026 NAME HERE contato@jao.dev.br
*/
package cmd

import (
	"fmt"
	"os"

	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/joaomortani/exevo-terra/internal/helpers"
)

var sharedAwsCfg aws.Config

var rootCmd = &cobra.Command{
	Use:   "exevo-terra",
	Short: "IaC Generator",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Name() == "help" || cmd.Name() == "version" {
			return nil
		}

		region, _ := cmd.Flags().GetString("region")
		profile, _ := cmd.Flags().GetString("profile")

		ctx := context.Background()

		cfg, err := helpers.LoadConfig(ctx, region, profile)
		if err != nil {
			log.Fatalf("Erro ao carregar AWS: %v", err)
		}
		sharedAwsCfg = cfg

		res, _ := cmd.Flags().GetString("resource")
		if res == "" {
			return fmt.Errorf("a flag --resource é obrigatória. Ex: --resource rds")
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetConfigName(".exevo-terra")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("resource", "t", "", "Resource to exevo-terra")
	rootCmd.PersistentFlags().StringP("profile", "p", "", "AWS Profile (SSO)")
	rootCmd.PersistentFlags().StringP("region", "r", "us-east-1", "AWS region")
}
