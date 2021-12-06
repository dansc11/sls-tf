package cmd

import (
	"fmt"
	"os"

	"github.com/dansc11/sls-tf/app"

	"github.com/spf13/cobra"
)

var serverlessYmlPath string

var rootCmd = &cobra.Command{
	Use:   "sls-tf",
	Short: "A serverless deployment framework using Terraform",
	Run: func(cmd *cobra.Command, args []string) {
		app.Plan(serverlessYmlPath)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&serverlessYmlPath, "sls-path", "serverless.yml", "location of the serverless.yml config file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
