package cmd

import (
  "fmt"
	"github.com/bwmarrin/discordgo"
)

func AllCommands() []*discordgo.ApplicationCommand {
  return []*discordgo.ApplicationCommand {
    NewPingCmd(),
  }
}

func RunCommand(session *discordgo.Session, interaction *discordgo.Interaction) error {
  err := Ping(session, interaction)
  if err != nil {
    return fmt.Errorf("ping: %w", err)
  }

  return nil
}
