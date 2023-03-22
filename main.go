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
	handler.Connect(*client)
}

func handlers(client *twitch.Client) {
	handler.Connected(client)
	handler.Logs(client)
	handler.Commands(client)
	handler.Alerts(client)
	// handler.Timers(client)
	handler.Connect(*client)
}
