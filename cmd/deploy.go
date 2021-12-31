package cmd

import (
	"github.com/dansc11/sls-tf/app"

	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the serverless resources using Terraform",
	Run: func(cmd *cobra.Command, args []string) {
		app.Deploy(workDir)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
