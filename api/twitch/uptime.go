package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	twitchBaseURL   = "https://api.twitch.tv/helix"
	twitchClientID  = "uefosio4dn1dwq1pde64p4sqx0uq1h"
	twitchChannelID = "197407231"
)

type Stream struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	StartedAt   time.Time `json:"started_at"`
	Type        string    `json:"type"`
	Language    string    `json:"language"`
	Title       string    `json:"title"`
	ViewerCount int       `json:"viewer_count"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/uptime", uptimeHandler)

	http.ListenAndServe(":8080", r)
}

func uptimeHandler(w http.ResponseWriter, r *http.Request) {
	apiUrl := fmt.Sprintf("%s/streams?user_id=%s", twitchBaseURL, twitchChannelID)

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Client-ID", twitchClientID)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data struct {
		Data []Stream `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(data.Data) == 0 {
		fmt.Fprint(w, "The channel is currently offline.")
	} else {
		stream := data.Data[0]
		uptime := time.Since(stream.StartedAt)

		fmt.Fprintf(w, "The channel has been live for %s.", uptime)
	}
}
