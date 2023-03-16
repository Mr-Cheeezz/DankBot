package handler

import (
	"fmt"
	"time"

	"github.com/TwiN/go-color"
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/mr-cheeezz/dankbot/env"
)

func init() {
	env.Load()
}

func Connected(client *twitch.Client) {
	DEBUG := env.Get("DEBUG")
	client.OnConnect(func() {
		client.Say(env.Get("CHANNEL_NAME"), "MrCheeezzBot | Golang version is now running. gopherDance ")
		if DEBUG == "true" {
			fmt.Println(color.InCyan("BOT STARTED IN DEBUG MODE"))
			time.Sleep(2 * time.Second)
			fmt.Println("")
		}
		fmt.Println(color.InWhite("Connected to channel"+":"+""), color.InBold(env.Get("CHANNEL_NAME")))
		fmt.Println("")
	})
}
