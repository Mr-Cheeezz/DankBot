package handler

import (
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/mr-cheeezz/dankbot/alerts"
)

func Alerts(client *twitch.Client) {
	alerts.Prime(client)
}
