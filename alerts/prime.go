package alerts

import (
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/mr-cheeezz/dankbot/env"
)

func init() {
	env.Load()
}

func Prime(client *twitch.Client) {
	// clientID := env.Get("CLIENT_ID")
	// clientSecret := env.Get("CLIENT_SECRET")
	// ChannelName := env.Get("CHANNEL_NAME")
	// ChannelID := env.Get("CHANNEL_ID")

	// apiClient, _ := helix.NewClient(&helix.Options{
	// 	ClientID:     clientID,
	// 	ClientSecret: clientSecret,
	// })

	// events := make(chan *helix.Event)
	// go apiClient.SubscribeToEvents(events, ChannelID, "subscribe")

	// for event := range events {
	// 	if event.Subscription.IsPrime {
	// 		msg := fmt.Sprintf("Thank you for subscribing with Prime, %s!", event.UserDisplayName)
	// 		client.Say(ChannelName, msg)
	// 	}
	// }

}
