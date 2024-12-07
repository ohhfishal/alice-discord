package handler

import (
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func (h *Handler) InteractionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	switch interaction.Type {
	case discordgo.InteractionApplicationCommand:
		err := h.Command.Run(session, interaction.Interaction)
		if err != nil {
			slog.Warn("running command: %w", err)
		}
	default:
		slog.Warn(fmt.Sprintf("unhandled interaction: %s", interaction.Type))
	}
}
