package windows

import (
	"github.com/epiclabs-io/winman"
	"github.com/ppodds/bot-cmd-manager/config"
	"github.com/rivo/tview"
)

func NewEditWindow(app *tview.Application, wm *winman.Manager, config *config.Config) *winman.WindowBase {
	w := wm.NewWindow().
		Show().
		SetTitle("Edit").
		Maximize()
	t := tview.NewBox()
	w.SetRoot(t).AddButton(&winman.Button{
		Symbol:  'X',
		OnClick: func() { app.Stop() },
	})
	return w
}
