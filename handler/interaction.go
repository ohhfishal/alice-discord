package handler

import (
	"fmt"
	"github.com/ohhfishal/alice-discord/cmd"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	switch interaction.Type {
	case discordgo.InteractionApplicationCommand:
		err := cmd.RunCommand(session, interaction.Interaction)
		if err != nil {
			slog.Warn("running command: %w", err)
		}
	default:
		slog.Warn(fmt.Sprintf("unhandled interaction: %s", interaction.Type))
	}
}
