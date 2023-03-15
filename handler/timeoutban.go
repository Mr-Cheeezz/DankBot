package handler

// import (
// 	"fmt"

// 	"github.com/TwiN/go-color"
// 	twitch "github.com/gempir/go-twitch-irc/v4"
// )

// func TimeoutBan(client *twitch.Client) {
// 	client.OnNotice(func(message twitch.Notice) {
// 		if message.ID == "timeout_success" {
// 			duration := message.Tags["ban-duration"]
// 			username := message.Tags["target-user-id"]

// 			fmt.Println(color.InRed(username, "has been timed out for", duration))

// 		} else if message.ID == "ban_success" {
// 			username := message.Tags["target-user-id"]

// 			fmt.Println(color.InRed(username, "has been banned"))

// 		}
// 	})

// }
