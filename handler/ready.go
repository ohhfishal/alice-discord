package handler

import (
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func OnReady(s *discordgo.Session, r *discordgo.Ready) {
	slog.Info(fmt.Sprintf("logged in as %s", r.User.String()))
}
