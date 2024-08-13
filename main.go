package main

import (
	"gitlab.com/ascenty/nemesis-cli/cmd"
	"gitlab.com/ascenty/nemesis-cli/config"
	"gitlab.com/ascenty/nemesis-cli/ui"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		return
	}

	app := ui.NewApp()
	cmd.Execute(app)
	if err := app.Run(); err != nil {
		panic(err)
	}

}
