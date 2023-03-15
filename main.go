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

func handlers(client *twitch.Client) {
	handler.Clear(client)
	handler.Connected(client)
	handler.Connect(client)
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		handler.Commands(client, message)
		handler.Logs(client)
	})
}
