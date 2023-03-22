package handler

import (
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/mr-cheeezz/dankbot/env"
	"github.com/mr-cheeezz/dankbot/timers"
)

func init() {
	env.Load()
}

func Timers(client *twitch.Client) {
	timers.Socials(client)
}
