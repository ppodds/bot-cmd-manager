package windows

import (
	"github.com/epiclabs-io/winman"
	"github.com/ppodds/bot-cmd-manager/config"
	"github.com/rivo/tview"
)

func NewConfigWindow(app *tview.Application, wm *winman.Manager, config *config.Config) *winman.WindowBase {
	w := wm.NewWindow().
		Show().
		SetDraggable(true).
		SetResizable(true).
		SetTitle("Config")
	form := tview.NewForm().AddPasswordField("Token", "", 0, '*', func(text string) {
		config.Token = text
	}).AddButton("Save", func() {
		config.Save()
		wm.RemoveWindow(w)
		wm.AddWindow(NewEditWindow(app, wm, config))
	})
	w.SetRoot(form).AddButton(&winman.Button{
		Symbol:  'X',
		OnClick: func() { 
			if (wm.WindowCount() == 1) {
				app.Stop()
			}
			wm.RemoveWindow(w)
		},
	})
	w.SetRect(0, 0, 30, 10)
	return w
}
