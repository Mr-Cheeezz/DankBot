package handler

import (
	"fmt"

	"github.com/TwiN/go-color"
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/mr-cheeezz/dankbot/env"
)

var excludedUsers = []string{"fossabot", "nightbot", "mr_cheeezzbot", "wizebot", "basementhelper"}

func Logs(client *twitch.Client) {
	messages(client)
	whispers(client)
}

func messages(client *twitch.Client) {
	env.Load()
	DEBUG := env.Get("DEBUG")
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if contains(excludedUsers, message.User.Name) {
			return
		}

		if DEBUG == "true" {
			fmt.Println((color.InGreen("#" + message.Channel + " " + "|")) + " " + (color.InPurple("[" + message.User.ID + "]")) + " " + (color.InBlue(message.User.Name + ":")) + " " + (color.InBlue(message.Message)) + " " + (color.InPurple("[" + message.ID + "]")))
		} else if DEBUG == "false" {
			fmt.Println("#"+message.Channel, "|"+" "+message.User.Name+":"+" "+message.Message)
		} else {
			fmt.Println(color.InRed("Failed to fetch message"))
		}
	})
}

func whispers(client *twitch.Client) {
	env.Load()
	DEBUG := env.Get("DEBUG")
	client.OnWhisperMessage(func(message twitch.WhisperMessage) {
		if DEBUG == "true" {
			fmt.Println((color.InGreen("whisper"+" | ") + (color.InPurple(("[") + message.User.ID + "]")) + " " + (color.InBlue(message.User.Name + ": " + message.Message))))
		} else if DEBUG == "false" {
			fmt.Println("whisper" + " | " + message.User.Name + " > " + env.Get("BOT_NAME") + ": " + message.Message)
		} else {
			fmt.Println(color.InRed("Failed to fetch whisper"))
		}
	})
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
