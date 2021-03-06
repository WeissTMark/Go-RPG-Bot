package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var UserList = map[string]Character{}
var classList = setClassDefaults()
var TOKEN = "NTU0MTc0OTc0OTgyODgxMjgw.XIShqg.JZZilNWo-sGI17MKNAUJuOqUvy0"

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	command := strings.Split(m.Content, " ")
	command[0] = strings.ToLower(command[0])

	switch command[0] {
	case "!botter":
		s.ChannelMessageSend(m.ChannelID, "Hello! How may I help you today?")
		break

	case "!random":
		s.ChannelMessageSend(m.ChannelID, "Sorry! This feature has not yet been implemented")
		break

	case "!level":
		level(s, m)
		break

	case "!newcharacter":
		newChar(s, m, command)
		break

	case "thesecretcode":
		ssshhh(s, m, command)
		break

	case "!setcolor":
		setColor(s, m, command)
		break

	case "!levels":
		levels(s, m)
		break

	case "!classes":
		listClasses(s, m)
		break
	}

}
