package handler

import (
	"fmt"
	"time"

	"github.com/TwiN/go-color"
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/mr-cheeezz/dankbot/env"
)

const lang = "Go"

func init() {
	env.Load()
}

func Connected(client *twitch.Client) {
	DEBUG := env.Get("DEBUG")
	client.OnConnect(func() {
		client.Say(env.Get("CHANNEL_NAME"), fmt.Sprintf("DankBot | %s version is now running. gopherDance", lang))
		if DEBUG == "true" {
			fmt.Println(color.InCyan("BOT STARTED IN DEBUG MODE"))
			time.Sleep(2 * time.Second)
			fmt.Println("")
		}

		fmt.Println(color.InWhite("Connected to channel"+":"+""), color.InBold(env.Get("CHANNEL_NAME")))
		fmt.Println("")
	})

}
