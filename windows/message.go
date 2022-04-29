package windows

import (
	"github.com/epiclabs-io/winman"
	"github.com/ppodds/bot-cmd-manager/app"
	"github.com/rivo/tview"
)

func NewMessageWindow(app *app.App, message string) *winman.WindowBase {
	w := app.Manager.NewWindow().
		Show().
		SetDraggable(true).
		SetResizable(true).
		SetTitle("Message")

	text := tview.NewTextView().SetText(message).SetTextAlign(tview.AlignCenter)

	buttonBar := tview.NewFlex().
		SetDirection(tview.FlexColumn)

	content := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(text, 0, 1, false).
		AddItem(buttonBar, 1, 0, true)

	buttonBar.AddItem(tview.NewButton("OK").SetSelectedFunc(func() {
		app.Manager.RemoveWindow(w)
		if app.Manager.WindowCount() == 0 {
			app.Stop()
		}
	}), 0, 1, true)

	w.SetRoot(content)

	w.SetRect(4, 2, 30, 6)

	app.SetFocus(w)

	return w
}
