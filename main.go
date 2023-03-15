package main

import (
	"github.com/TwiN/go-color"
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/mr-cheeezz/dankbot/env"
	"github.com/mr-cheeezz/dankbot/handler"
)

func main() {
	env.Load()

	var (
		OAuth   = env.Get("BOT_OAUTH")
		BotName = env.Get("BOT_NAME")

		ChannelName = env.Get("CHANNEL_NAME")
	)

	client := twitch.NewClient(BotName, "oauth:"+OAuth)
	client.Join(ChannelName)

	handler.Clear(client)
	handler.Connected(client)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		handler.Commands(client, message)
		handler.Logs(client)
	})

	err := client.Connect()
	if err != nil {
		panic(color.InRed(err))
	}
}
