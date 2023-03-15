package handler

import (
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/mr-cheeezz/dankbot/commands"
)

func Commands(client *twitch.Client, message twitch.PrivateMessage) {
	if message.Message == "!version" {
		commands.Version(client, &message)
	}
}
