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
	guildID  string
	commands []*discordgo.ApplicationCommand
	viewing  *discordgo.ApplicationCommand
	detail   *tview.Flex
}

// create a new EditWindow
// if guildID is empty string, load global command instead
// return nil if get some error
func NewEditWindow(app *app.App, guildID string) *EditWindow {
	if app.Session == nil {
		s, err := discordgo.New("Bot " + app.Config.Token)
		if err != nil {
			panic(err)
		}
		app.Session = s
	}

	t := app.Manager.NewWindow()
	w := (&EditWindow{
		WindowBase: *t,
		guildID:    guildID,
	})
	app.RemoveWindow(t)
	app.AddWindow(w)
	w.WindowBase.
		Show().
		SetTitle("Edit").
		SetResizable(true).
		SetDraggable(true)
	c := tview.NewFlex().SetDirection(tview.FlexRow)
	w.SetRoot(c).AddButton(&winman.Button{
		Symbol: 'X',
		OnClick: func() {
			app.Manager.RemoveWindow(w)
			if app.Manager.WindowCount() == 0 {
				app.Stop()
			}
		},
	})

	commands, err := app.Session.ApplicationCommands(app.Config.ApplicationID, w.guildID)
	if err != nil {
		if err == discordgo.ErrUnauthorized {
			NewMessageWindow(app, "Can't get commands! Please check your bot token, application ID.")
		} else {
			NewMessageWindow(app, "Can't get commands! Please check your guild ID.")
		}
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
	if w.guildID == "" {
		c.AddItem(tview.NewButton("Config").SetSelectedFunc(func() {
			app.SetFocus(NewConfigWindow(app))
		}), 1, 1, false)
		c.AddItem(tview.NewButton("View Guild Commands").SetSelectedFunc(func() {
			app.SetFocus(NewGuildWindow(app))
		}), 1, 1, false)
	}

	c.AddItem(d, 1, 1, false)
	c.AddItem(w.detail, 0, 1, false)
	c.AddItem(tview.NewButton("Delete").SetSelectedFunc(func() {
		if w.viewing == nil {
			NewMessageWindow(app, "Please select a command first.")
			return
		}
		err := app.Session.ApplicationCommandDelete(app.Config.ApplicationID, w.guildID, w.viewing.ID)
		if err != nil {
			NewMessageWindow(app, "Can't delete command! Please check your bot token, application ID.")
			return
		}
		NewMessageWindow(app, "Command deleted.")
		i, _ := d.GetCurrentOption()
		w.commands = append(w.commands[:i], w.commands[i+1:]...)
		d.RemoveOption(i)
		i, _ = d.GetCurrentOption()
		if i == -1 {
			w.viewing = nil
		} else {
			w.viewing = w.commands[i]
		}
		w.ChangeViewingCommand()
	}), 1, 1, false)

	w.SetRect(0, 0, 50, len(w.commands)+5)
	return w
}

func (w *EditWindow) ChangeViewingCommand() {
	w.detail.Clear()
	if w.viewing == nil {
		w.detail.AddItem(tview.NewTextView().SetText("No command selected."), 0, 1, false)
	} else {
		w.detail.AddItem(tview.NewTextView().SetText("ID: "+w.viewing.ID), 1, 1, false)
		w.detail.AddItem(tview.NewTextView().SetText("Application ID: "+w.viewing.ApplicationID), 1, 1, false)
		w.detail.AddItem(tview.NewTextView().SetText("Version: "+w.viewing.Version), 1, 1, false)
		w.detail.AddItem(tview.NewTextView().SetText(fmt.Sprintf("Default Permission: %t", *w.viewing.DefaultPermission)), 1, 1, false)
		w.detail.AddItem(tview.NewTextView().SetText("Type: "+fmtCmdType(w.viewing.Type)), 1, 1, false)
		w.detail.AddItem(tview.NewTextView().SetText("Name: "+w.viewing.Name), 1, 1, false)
		w.detail.AddItem(tview.NewTextView().SetText("Description: "+w.viewing.Description), 0, 1, false)
	}
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
