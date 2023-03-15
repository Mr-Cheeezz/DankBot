package roblox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type UserPresence struct {
	UserPresenceType int    `json:"userPresenceType"`
	LastLocation     string `json:"lastLocation"`
	PlaceId          int    `json:"placeId"`
	UniverseId       int    `json:"universeId"`
}

type UserPresencesResponse struct {
	UserPresences []UserPresence `json:"userPresences"`
}

func GetPresence(userId int, cookie string) (*UserPresence, error) {
	url := "https://presence.roblox.com/v1/presence/users"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{"userIds": [%d]}`, userId))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", fmt.Sprintf(".ROBLOSECURITY=%s", cookie))
	csrfToken, err := GetXcsrf(cookie)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-CSRF-TOKEN", csrfToken)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response UserPresencesResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	if len(response.UserPresences) > 0 {
		return &response.UserPresences[0], nil
	} else {
		return nil, fmt.Errorf("user %d not found", userId)
	}
}

func GetGameName(userId int, cookie string) (string, error) {
	presence, err := GetPresence(userId, cookie)
	if err != nil {
		return "", err
	}

	if presence.UserPresenceType != 2 {
		return "", fmt.Errorf("user %d is not playing a game", userId)
	}

	url := fmt.Sprintf("https://games.roblox.com/v1/games/multiget-place-details?placeIds=%d", presence.PlaceId)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	data, ok := response[strconv.Itoa(presence.PlaceId)].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("failed to parse game data")
	}

	name, ok := data["name"].(string)
	if !ok {
		return "", fmt.Errorf("failed to parse game name")
	}

	return name, nil
}
