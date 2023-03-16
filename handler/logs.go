package handler

import (
	"github.com/gempir/go-twitch-irc/v4"
	log "github.com/mr-cheeezz/dankbot/handler/functions"
)

func Logs(client *twitch.Client) {
	log.Messages(client)
	log.Whispers(client)
	log.Clear(client)
}
