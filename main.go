package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func main() {
	session, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		slog.Error(fmt.Sprintf("starting bot: %w", err))
		os.Exit(1)
	}

	// Log ready
	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		slog.Info(fmt.Sprintf("logged in as %s", r.User.String()))
	})

	// First command!
	session.AddHandler(func(s *discordgo.Session, i *discordgo.MessageCreate) {
		fmt.Println("CALLING HANDLER")
		fmt.Println("DONE")
	})

	var commands = []*discordgo.ApplicationCommand{{Name: "echo", Description: "Say something through a bot", Options: []*discordgo.ApplicationCommandOption{{Name: "message", Description: "Contents of the message", Type: discordgo.ApplicationCommandOptionString, Required: true}, {Name: "author", Description: "Whether to prepend message's author", Type: discordgo.ApplicationCommandOptionBoolean}}}}

	app := os.Getenv("APP_ID")
	guild := ""
	_, err = session.ApplicationCommandBulkOverwrite(app, guild, commands)
	if err != nil {
		slog.Error(fmt.Sprintf("registering commands: %s", err))
    return
	}

	err = session.Open()
	if err != nil {
		slog.Error(fmt.Sprintf("starting session: %w", err))
		os.Exit(1)
	}
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	<-sigch

	err = session.Close()
	if err != nil {
		slog.Error(fmt.Sprintf("could not close gracefully: %w", err))
	}

}
