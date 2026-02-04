package rds

import (
	"github.com/joaomortani/exevo-terra/internal/adapter"
	"github.com/joaomortani/exevo-terra/internal/generator"
	"github.com/joaomortani/exevo-terra/internal/provider"
	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Preview dos blocos de código Terraform",
	RunE: func(cmd *cobra.Command, args []string) error {
		instances, err := provider.FetchRDSInstances(cmd.Context(), sharedAwsCfg)
		if err != nil {
			return err
		}

		cleanList := adapter.BatchAwsToDomain(instances)
		return generator.PlanHCL(cleanList)
	},
}

func init() {
	RdsCmd.AddCommand(planCmd)
}
