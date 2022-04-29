package windows

import (
	"github.com/epiclabs-io/winman"
	"github.com/ppodds/bot-cmd-manager/app"
	"github.com/rivo/tview"
)

func NewGuildWindow(app *app.App) *winman.WindowBase {
	w := app.Manager.NewWindow().
		Show().
		SetDraggable(true).
		SetResizable(true).
		SetTitle("Guild Command")

	guildID := ""

	form := tview.NewForm().AddInputField("Guild ID", "", 0, nil, func(text string) {
		guildID = text
	}).AddButton("Next", func() {
		if guildID != "" {
			app.Manager.AddWindow(NewEditWindow(app, guildID))
			app.Manager.RemoveWindow(w)
		}
	})

	w.SetRoot(form).AddButton(&winman.Button{
		Symbol: 'X',
		OnClick: func() {
			app.Manager.RemoveWindow(w)
		},
	})

	w.SetRect(0, 0, 30, 10)

	return w
}
