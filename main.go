package main

import (
	"fmt"
	"os"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	OAuth := os.Getenv("OAUTH")
	BotName := os.Getenv("BOT_NAME")
	ChannelName := os.Getenv("CHANNEL_NAME")

	client := twitch.NewClient(BotName, "oauth:"+OAuth)
	client.Join(ChannelName)

	client.OnConnect(func() {
		client.Say(ChannelName, fmt.Sprintf("MrCheeezzBot | Golang version is now running. gopherDance "))
		fmt.Println("Bot connected to channel: ", ChannelName)
	})

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.User.Name + ": " + message.Message)
	})

	err := client.Connect()
	if err != nil {
		panic(err)
	}

}
