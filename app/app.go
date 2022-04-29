package app

import (
	"github.com/bwmarrin/discordgo"
	"github.com/epiclabs-io/winman"
	"github.com/ppodds/bot-cmd-manager/config"
	"github.com/rivo/tview"
)

type App struct {
	*tview.Application
	*winman.Manager
	*config.Config
	*discordgo.Session
}

func NewApp() *App {
	app := &App{
		Application: tview.NewApplication(),
		Manager:     winman.NewWindowManager(),
	}
	return app
}