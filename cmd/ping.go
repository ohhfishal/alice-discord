package cmd

import (
	"github.com/bwmarrin/discordgo"
)

func NewPingCmd() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "reply to a ping",
	}
}

func Ping(session *discordgo.Session, interaction *discordgo.Interaction) error {
	return session.InteractionRespond(interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "pong",
		},
	})
}
