package api

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthResponse struct {
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	Expires      string `json:"expires"`
}

func Login(baseURL, auth string) (*AuthResponse, error) {
	url := fmt.Sprintf("%s/v1/tokens?grant_type=password", baseURL)

	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", auth)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var authResponse AuthResponse

	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return nil, err
	}

	return &authResponse, nil
}
