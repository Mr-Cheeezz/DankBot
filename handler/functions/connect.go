package functions

import (
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/mr-cheeezz/dankbot/env"
)

func init() {
	env.Load()
}

func Join(client *twitch.Client) {
	client.Join(env.Get("CHANNEL_NAME"))
}

func Connect(client *twitch.Client) {
	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
