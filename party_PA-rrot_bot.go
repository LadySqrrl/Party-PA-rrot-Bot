package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func main() {

	var token string

	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("could not begin discord session", err)
	}

}
