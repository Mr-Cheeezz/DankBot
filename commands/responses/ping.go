package commands

import (
	"fmt"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

func Ping(client *twitch.Client, message *twitch.PrivateMessage) {
	client.Say(message.Channel, "Pong! FeelsBingMan")
	time.Sleep(1 * time.Second)
	client.Reply(message.Channel, message.ID, fmt.Sprintf("@%s, your ping is %s.", message.User.Name, time.Since(message.Time)))
}
