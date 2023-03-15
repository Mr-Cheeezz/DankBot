package handler

import (
	"fmt"

	"github.com/TwiN/go-color"
	"github.com/gempir/go-twitch-irc/v4"
)

func Clear(client *twitch.Client) {
	client.OnClearChatMessage(func(message twitch.ClearChatMessage) {
		if message.Message == "" {
			fmt.Println(color.InYellow("The chat has been cleard"))
		}
	})
}
