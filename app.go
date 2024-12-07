package main

import (
	"github.com/ohhfishal/alice-discord/handler"

	"github.com/bwmarrin/discordgo"
)

type App interface {
	Open() error
	Close() error
}

func NewApp(getenv func(string) string) (App, error) {
	session, err := discordgo.New("Bot " + getenv("TOKEN"))
	if err != nil {
		return nil, err
	}

	handler := handler.NewHandler(getenv)
	handler.RegisterSession(session)

	return session, nil
}
