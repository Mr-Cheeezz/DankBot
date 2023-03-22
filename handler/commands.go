package handler

import (
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/mr-cheeezz/dankbot/commands"
)

func Commands(client *twitch.Client) {
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if message.Message == "!version" {
			commands.Version(client, &message)
		}

		if message.Message == "!ping" {
			commands.Ping(client, &message)
		}
	})
}
