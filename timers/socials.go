package timers

import (
	"math/rand"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/mr-cheeezz/dankbot/env"
)

func Socials(client *twitch.Client) {
	env.Load()
	ticker := time.NewTicker(time.Minute * 1)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			promo := []string{
				"!discord",
				"!youtube",
				"!facebook",
			}
			rand.Seed(time.Now().UnixNano())
			promoTimer := promo[rand.Intn(len(promo))]
			client.Say(env.Get("CHANNEL_NAME"), promoTimer)
		}
	}
}
