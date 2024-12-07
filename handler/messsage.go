package handler

import (
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func (h *Handler) MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	return
	if session.State.User.ID == message.Author.ID {
		slog.Info(fmt.Sprintf("replied: [%s] %s", message.Author, message.Content))
		return
	}
	slog.Info(fmt.Sprintf("message: [%s] %s", message.Author, message.Content))
}
