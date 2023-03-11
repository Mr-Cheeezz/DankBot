package roblox

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func getGame(userid string, cookie string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://presence.roblox.com/v1/presence/users", nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Cookie", cookie)

	q := req.URL.Query()
	q.Add("userIds", userid)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var presence struct {
		UserPresences []struct {
			UserPresenceStatusType int `json:"userPresenceType"`
			PlaceID                int `json:"placeId"`
			GameID                 int `json:"gameId"`
		} `json:"userPresences"`
	}

	err = json.NewDecoder(resp.Body).Decode(&presence)
	if err != nil {
		return "", err
	}

	if len(presence.UserPresences) == 0 {
		return "", errors.New("User is not currently playing any game")
	}

	gameid := presence.UserPresences[0].GameID

	resp, err = client.Get(fmt.Sprintf("https://games.roblox.com/v1/games/%d", gameid))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var game commands.Game
	err = json.NewDecoder(resp.Body).Decode(&game)
	if err != nil {
		return "", err
	}

	return game.Name, nil
}
