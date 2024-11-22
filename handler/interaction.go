package handler

import (
  "fmt"
  "log/slog"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
  slog.Info(fmt.Sprintf("Interaction: %s", interaction.Data))

  err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
      Type: discordgo.InteractionResponseChannelMessageWithSource,
      Data: &discordgo.InteractionResponseData{
        Content: "ligma balls",
      },
  })

  if err != nil {
    slog.Error(fmt.Sprintf("responding to interaction: %w", err))
  }
}
