package twitch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mr-cheeezz/dankbot/env"
)

func init() {
	env.Load()
}

func IsLive() (bool, error) {

	var (
		botOAuth      = env.Get("BOT_OAUTH")
		channelName   = env.Get("CHANNEL_NAME") // pokimane
		clientID      = "kimne78kx3ncx6brgo4mv6wki5h1ko"
		gqlEndpoint   = "https://gql.twitch.tv/gql"
		queryTemplate = `[{"operationName":"VideoPlayerStreamInfoOverlayChannel","variables":{"channel":"%s"},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"a5f2e34d626a9f4f5c0204f910bab2194948a9502089be558bb6e779a9e1b3d2"}}}]`
	)

	body := fmt.Sprintf(queryTemplate, channelName)
	req, err := http.NewRequest("POST", gqlEndpoint, strings.NewReader(body))
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", "OAuth "+botOAuth)
	req.Header.Set("Client-ID", clientID)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var data []struct {
		Data struct {
			User struct {
				Stream *struct{} `json:"stream"`
			} `json:"user"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return false, err
	}

	return data[0].Data.User.Stream != nil, nil
}
