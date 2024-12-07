package handler

import (
	"fmt"
	"github.com/ohhfishal/alice-discord/cmd"
	"github.com/ohhfishal/alice-discord/config"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	Config  config.Config
	Command cmd.Handler
}

func (h *Handler) RegisterSession(session *discordgo.Session) error {
	newCmds, err := session.ApplicationCommandBulkOverwrite(h.Config.AppID, "", cmd.AllCommands())
	if err != nil {
		return err
	}
	for _, command := range newCmds {
		slog.Debug(fmt.Sprintf("registered command: %s", command.Name))
	}

	session.AddHandler(h.OnReady)
	session.AddHandler(h.MessageCreate)
	session.AddHandler(h.InteractionCreate)
	return nil
}

func NewHandler(getenv func(string) string) *Handler {
	cfg := config.NewConfig(getenv)
	handler := &Handler{
		Config: cfg,
		Command: cmd.Handler{
			Config: cfg,
		},
	}
	return handler
}
