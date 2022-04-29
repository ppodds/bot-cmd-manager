package main

import (
	"github.com/epiclabs-io/winman"
	"github.com/ppodds/bot-cmd-manager/config"
	"github.com/ppodds/bot-cmd-manager/windows"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	wm := winman.NewWindowManager()
	config := config.NewConfig()

	if !config.Load() {
		wm.AddWindow(windows.NewConfigWindow(app, wm, config))
	} else {
		wm.AddWindow(windows.NewEditWindow(app, wm, config))
	}

	if err := app.SetRoot(wm, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
