package main

import (
  "fmt"
  "log/slog"
  "github.com/ohhfishal/alice-discord/cmd"
  "github.com/ohhfishal/alice-discord/handler"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
  AppID string
}

type App struct {
  Session *discordgo.Session
  Config Config

}

func NewConfig(getenv func(string)string) Config {
  return Config {
    AppID: getenv("APP_ID"),
  }
}

func NewApp(getenv func(string)string) (*App, error) {
  var app App
	session, err := discordgo.New("Bot " + getenv("TOKEN"))
	if err != nil {
    return nil, err
	}

  config := NewConfig(getenv)

  app.Session = session
  app.Config = config
  app.setupHandlers()

  commands := cmd.AllCommands()

  newCmds, err := app.Session.ApplicationCommandBulkOverwrite(app.Config.AppID, "", commands)
  if err != nil {
    return nil, err
  }
  for _, command := range newCmds {
    slog.Info(fmt.Sprintf("registered command: %s", command.Name))
  }

  return &app, nil
}


func (app *App) Start() error {
	return app.Session.Open()
}

func (app *App) Shutdown() error {
	return app.Session.Close()
}

func (app *App) registerCommands() {

}

func (app *App) setupHandlers() {
	app.Session.AddHandler(handler.OnReady)
	app.Session.AddHandler(handler.MessageCreate)
	app.Session.AddHandler(handler.InteractionCreate)
}
