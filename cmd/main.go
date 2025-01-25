package main

import (
	"discord_go/config"
	"discord_go/events"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := config.GetToken()

	dg, err := discordgo.New("Bot" + token)
	if err != nil {
		fmt.Println("Error creating Discord Session", err)
		return
	}

	dg.AddHandler(events.OnMessageCreate)

	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}
