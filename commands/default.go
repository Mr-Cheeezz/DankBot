package commands

import (
	"time"

	"github.com/Mr-Cheeezz/DankBot/utils"
	"github.com/gempir/go-twitch-irc/v4"
)

var cooldowns = make(map[string]*utils.CommandCooldown)

func init() {
	cooldowns["hello"] = &utils.CommandCooldown{}
}

func HelloWorld(client *twitch.Client, message twitch.PrivateMessage) {
	cooldown := 30 * time.Second
	command := "hello"

	if !cooldowns[command].Check(command, cooldown) {
		return
	}

	cooldowns[command].Update(command)

	client.Reply(message.Channel, message.ID, "Hello!")
}
