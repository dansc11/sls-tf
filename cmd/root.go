package cmd

import (
	"fmt"
	"os"

	"github.com/dansc11/sls-tf/app"

	"github.com/spf13/cobra"
)

var workDir string

var rootCmd = &cobra.Command{
	Use:   "sls-tf",
	Short: "A serverless deployment framework using Terraform",
	Run: func(cmd *cobra.Command, args []string) {
		app.Plan(workDir)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&workDir, "sls-path", "p", "serverless.yml", "location of the serverless config files")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
