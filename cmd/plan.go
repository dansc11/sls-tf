package cmd

import (
	"github.com/dansc11/sls-tf/app"

	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "A serverless deployment framework using Terraform",
	Run: func(cmd *cobra.Command, args []string) {
		app.Plan(workDir)
	},
}

func init() {
	rootCmd.AddCommand(planCmd)
}
