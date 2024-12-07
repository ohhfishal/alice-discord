package cmd

import (
	"github.com/bwmarrin/discordgo"
)

func NewTaskCmd() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "task",
		Description: "create a new task",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "description",
				Description: "Description for the task.",
				Required:    true,
				Type:        discordgo.ApplicationCommandOptionString,
			},
		},
	}
}

func (h *Handler) Task(session *discordgo.Session, interaction *discordgo.Interaction) error {
	return session.InteractionRespond(interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "task create",
		},
	})
}
