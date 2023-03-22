package main

import (
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/mr-cheeezz/dankbot/env"
	"github.com/mr-cheeezz/dankbot/handler"
)

func init() {
	env.Load()
}

func main() {
	client := twitch.NewClient(env.Get("BOT_NAME"), "oauth:"+env.Get("BOT_OAUTH"))

	handlers(client)
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

func handlers(client *twitch.Client) {
	handler.Connected(client)
	handler.Logs(client)
	handler.Commands(client)
	handler.Alerts(client)
	handler.Connect(*client)
}
