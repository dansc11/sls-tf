package cmd

import (
	"github.com/dansc11/sls-tf/app"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new sls-tf project",
	Run: func(cmd *cobra.Command, args []string) {
		app.InitProject()
	},
}
