package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var Token string = "YOUR_BOT_TOKEN" // Replace this with your bot token

// This function will be called when the bot is ready
func ready(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("Bot is now online!")
}

// This function will handle incoming messages
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Respond to the "!hello" command
	if m.Content == "!hello" {
		s.ChannelMessageSend(m.ChannelID, "Hello, I am your bot!")
	}
}

func main() {
	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the ready and messageCreate functions
	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening for events
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until the program is terminated (Ctrl+C)
	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}
