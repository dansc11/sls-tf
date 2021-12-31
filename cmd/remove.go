package cmd

import (
	"github.com/dansc11/sls-tf/app"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the serverless resources using Terraform",
	Run: func(cmd *cobra.Command, args []string) {
		app.Remove(workDir)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
