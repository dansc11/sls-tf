package cmd

import (
	"github.com/dansc11/sls-tf/app"
	"github.com/spf13/cobra"
)

var projectPath string

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&projectPath, "path", "p", "", "destination path for the project")
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new sls-tf project",
	Run: func(cmd *cobra.Command, args []string) {
		app.InitProject(projectPath)
	},
}
