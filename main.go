package main

import (
	"github.com/ppodds/bot-cmd-manager/app"
	"github.com/ppodds/bot-cmd-manager/config"
	"github.com/ppodds/bot-cmd-manager/windows"
)

func main() {
	app := app.NewApp()
	app.Config = config.NewConfig()

	if !app.Config.Load() {
		app.Manager.AddWindow(windows.NewConfigWindow(app))
	} else {
		app.Manager.AddWindow(windows.NewEditWindow(app, ""))
	}

	if err := app.SetRoot(app.Manager, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
