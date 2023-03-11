package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/joho/godotenv"
)

type Game struct {
	Name string `json:"name"`
}

func HandleRobloxGameCommand(client *twitch.Client, message twitch.PrivateMessage) {
	godotenv.Load(".env")
	// Get user ID from Twitch display name
	userID := 129152945
	Cookie := os.Getenv("ROBLOX_TOKEN")

	// Get current game from Roblox Presence API
	gameName, err := getCurrentGameName(userID, Cookie)

	if err != nil {
		client.Reply(message.Channel, message.ID, "Error getting game information.")
		return
	}

	// Send game name to chat
	client.Reply(message.Channel, message.ID, fmt.Sprintf("The streamer is currently playing %s", gameName))
}

func getCurrentGameName(userID int, cookie string) (string, error) {
	// Send GET request to Roblox Presence API
	req, err := http.NewRequest("GET", fmt.Sprintf("https://presence.roblox.com/v1/presence/users/%d/games?sortOrder=Asc&limit=1", userID), nil)

	if err != nil {
		return "", err
	}

	// Set cookie header for authentication
	req.Header.Set("Cookie", cookie)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	// Parse response JSON
	var games []Game
	err = json.NewDecoder(resp.Body).Decode(&games)

	if err != nil {
		return "", err
	}

	// Return game name
	if len(games) > 0 {
		return games[0].Name, nil
	} else {
		return "", fmt.Errorf("no games found for user ID %d", userID)
	}
}
