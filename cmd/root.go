/*
Copyright © 2026 NAME HERE contato@jao.dev.br
*/
package cmd

import (
	"os"

	"github.com/joaomortani/exevo-terra/cmd/rds"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "exevo-terra",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	rootCmd.AddCommand(rds.RdsCmd)

	rootCmd.PersistentFlags().StringP("profile", "p", "", "AWS Profile (SSO)")
	rootCmd.PersistentFlags().StringP("region", "r", "us-east-1", "AWS region")
}
