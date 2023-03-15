package commands

import (
	"github.com/gempir/go-twitch-irc/v4"
)

func Version(client *twitch.Client, message *twitch.PrivateMessage) {
	client.Reply(message.Channel, message.ID, "MrCheeezzBot - Golang | 1.0")
}
