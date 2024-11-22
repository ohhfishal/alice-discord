package main

import (
  "github.com/ohhfishal/alice-discord/handler"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
}

type App struct {
  Session *discordgo.Session
  Config Config

}

func NewConfig(getenv func(string)string) Config {
  return Config {
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
  return &app, nil
}

func (app *App) Start() error {
	return app.Session.Open()
}

func (app *App) Shutdown() error {
	return app.Session.Close()
}

func (app *App) setupHandlers() {
	app.Session.AddHandler(handler.OnReady)
	app.Session.AddHandler(handler.MessageCreate)
	app.Session.AddHandler(handler.InteractionCreate)
}
