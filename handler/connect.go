package handler

import (
	"github.com/TwiN/go-color"
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/mr-cheeezz/dankbot/env"
)

func init() {
	env.Load()
}

func Connect(client *twitch.Client) {
	client.Join(env.Get("CHANNEL_NAME"))

	err := client.Connect()
	if err != nil {
		panic(color.InRed(err))
	}
}
