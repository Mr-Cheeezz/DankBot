package roblox

import (
	"fmt"
	"net/http"
)

func GetXcsrf(cookie string) (string, error) {
	const url = "https://auth.roblox.com/v2/logout"

	if cookie == "" {
		return "", fmt.Errorf("cookie is required")
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Cookie", fmt.Sprintf(".ROBLOSECURITY=%s", cookie))

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}

	csrfToken := resp.Header.Get("x-csrf-token")
	if csrfToken == "" {
		return "", fmt.Errorf("csrf token not found")
	}

	return csrfToken, nil
}
