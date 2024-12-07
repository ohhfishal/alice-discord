package cmd

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/ohhfishal/alice-discord/config"
)

type Handler struct {
	Config config.Config
}

func AllCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		NewPingCmd(),
		NewTaskCmd(),
	}
}

func (h *Handler) Run(session *discordgo.Session, interaction *discordgo.Interaction) error {
	name := interaction.ApplicationCommandData().Name
	switch name {
	case "ping":
		return h.Ping(session, interaction)
	case "task":
		return h.Task(session, interaction)
	default:
		return fmt.Errorf("unknown command: %s", name)
	}
}
