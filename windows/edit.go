package windows

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/epiclabs-io/winman"
	"github.com/ppodds/bot-cmd-manager/app"
	"github.com/rivo/tview"
)

type EditWindow struct {
	winman.WindowBase
	guildID string
	commands []*discordgo.ApplicationCommand
	viewing *discordgo.ApplicationCommand
	detail *tview.Flex
}

// create a new EditWindow
// if guildID is empty string, load global command instead
func NewEditWindow(app *app.App, guildID string) *EditWindow {
	if app.Session == nil {
		s, err := discordgo.New("Bot " + app.Config.Token)
		if (err != nil) {
			panic(err)
		}
		app.Session = s
	}
	
	w := (&EditWindow{
		WindowBase: *app.Manager.NewWindow(),
		guildID: guildID,
	})
	w.WindowBase.
		Show().
		SetTitle("Edit").
		Maximize()
	c := tview.NewFlex().SetDirection(tview.FlexRow)
	w.SetRoot(c).AddButton(&winman.Button{
		Symbol:  'X',
		OnClick: func() { app.Stop() },
	})

	commands, err := app.Session.ApplicationCommands(app.Config.ApplicationID, w.guildID)
	if (err != nil) {
		panic(err)
	}

	w.commands = commands

	d := tview.NewDropDown().SetLabel("Command: ")
	d.SetLabelColor(tview.Styles.PrimaryTextColor)

	for _, command := range commands {
		// https://stackoverflow.com/questions/7053965/when-using-callbacks-inside-a-loop-in-javascript-is-there-any-way-to-save-a-var
		local := command
		d.AddOption(command.Name, func() {
			w.viewing = local
			w.ChangeViewingCommand()
		})
	}
	
	w.detail = tview.NewFlex().SetDirection(tview.FlexRow)

	c.AddItem(d, 1, 1, false)
	c.AddItem(w.detail, 0, 1, false)
	c.AddItem(tview.NewButton("Delete"), 1, 1, false)
	return w
}

func (w *EditWindow) ChangeViewingCommand() {
	w.detail.Clear()
	w.detail.AddItem(tview.NewTextView().SetText("ID: " + w.viewing.ID), 1, 1, false)
	w.detail.AddItem(tview.NewTextView().SetText("Application ID: " + w.viewing.ApplicationID), 1, 1, false)
	w.detail.AddItem(tview.NewTextView().SetText("Version: " + w.viewing.Version), 1, 1, false)
	w.detail.AddItem(tview.NewTextView().SetText(fmt.Sprintf("Default Permission: %t", *w.viewing.DefaultPermission)), 1, 1, false)
	w.detail.AddItem(tview.NewTextView().SetText("Type: " + fmtCmdType(w.viewing.Type)), 1, 1, false)
	w.detail.AddItem(tview.NewTextView().SetText("Name: " + w.viewing.Name), 1, 1, false)
	w.detail.AddItem(tview.NewTextView().SetText("Description: " + w.viewing.Description), 0, 1, false)
}

func fmtCmdType(cmdType discordgo.ApplicationCommandType) string {
	switch cmdType {
	case discordgo.ChatApplicationCommand:
		return "Slash Command"
	case discordgo.UserApplicationCommand:
		return "User"
	case discordgo.MessageApplicationCommand:
		return "Message"
	default:
		return "Unknown"
	}
}
