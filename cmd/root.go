package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/ascenty/nemesis-cli/ui"
)

var rootCmd = &cobra.Command{
	Use:   "nemesis-cli",
	Short: "Nemesis CLI",
	Long:  `Nemesis CLI`,
}

func Execute(app *ui.App) {
	rootCmd.AddCommand(loginCmd)
	app.SetRootCommand(rootCmd)
}
