package handler

import (
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/mr-cheeezz/dankbot/env"
)

func Connect(client twitch.Client) {
	join(&client)
	connect(&client)
}

func join(client *twitch.Client) {
	client.Join(env.Get("CHANNEL_NAME"))
}

func connect(client *twitch.Client) {
	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
