package handler

import (
  "fmt"
  "log/slog"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
  slog.Info(fmt.Sprintf("Receieved: [%s] %s", message.Author, message.Content))
}
