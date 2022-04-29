package windows

import (
	"github.com/epiclabs-io/winman"
	"github.com/ppodds/bot-cmd-manager/app"
	"github.com/rivo/tview"
)

func NewConfigWindow(app *app.App) *winman.WindowBase {
	w := app.Manager.NewWindow().
		Show().
		SetDraggable(true).
		SetResizable(true).
		SetTitle("Config")
	form := tview.NewForm().AddPasswordField("Token", app.Config.Token, 0, '*', func(text string) {
		app.Config.Token = text
	}).AddInputField("Application ID", app.Config.ApplicationID, 0, nil, func(text string) {
		app.Config.ApplicationID = text
	}).AddButton("Save", func() {
		app.Config.Save()
		app.Manager.RemoveWindow(w)
		app.Manager.AddWindow(NewEditWindow(app, ""))
	})
	w.SetRoot(form).AddButton(&winman.Button{
		Symbol: 'X',
		OnClick: func() {
			app.Manager.RemoveWindow(w)
			if app.Manager.WindowCount() == 0 {
				app.Stop()
			}
		},
	})
	w.SetRect(0, 0, 100, 10)
	return w
}
