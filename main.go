package main

import (
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/mr-cheeezz/dankbot/env"
	"github.com/mr-cheeezz/dankbot/handler"
	"github.com/mr-cheeezz/dankbot/web"
)

func init() {
	env.Load()
}

func main() {
	client := twitch.NewClient(env.Get("BOT_NAME"), "oauth:"+env.Get("BOT_OAUTH"))

	handlers(client)
	web.Start()
}

func handlers(client *twitch.Client) {
	handler.Connected(client)
	handler.Logs(client)
	handler.Connect(client)
	handler.Commands(client)
}
