package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   uint   `json:"expires_in"`
	Sub         string `json:"sub"`
}

type AuctionDataResponse struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Auctions []struct {
		ID   int `json:"id"`
		Item struct {
			ID int `json:"id"`
		} `json:"item"`
		Quantity  int    `json:"quantity"`
		UnitPrice int    `json:"unit_price"`
		TimeLeft  string `json:"time_left"`
	} `json:"auctions"`
}

func GetAuctionData(token string, last_request_time time.Time) (auction_data *AuctionDataResponse, new_data bool) {
	c := http.Client{Timeout: time.Duration(20) * time.Second}
	req, _ := http.NewRequest("GET", "https://us.api.blizzard.com/data/wow/auctions/commodities?namespace=dynamic-us&locale=en_US", nil)

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("If-Modified-Since", last_request_time.Format(time.RFC1123))

	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Error occurred %s", err)
		return nil, true
	}
	if resp.StatusCode == 304 {
		return nil, false
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Recieved bad status code: %d", resp.StatusCode)
		return nil, true
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if body != nil {
		var responseData AuctionDataResponse
		err := json.Unmarshal(body, &responseData)
		if err != nil {
			fmt.Printf("Unmarshal error %s", err)
			return nil, true
		}
		return &responseData, true
	}
	return nil, true
}

func GetToken(client_id string, client_secret string) *AuthResponse {
	c := http.Client{Timeout: time.Duration(20) * time.Second}
	formData := url.Values{
		"grant_type": {"client_credentials"},
	}
	req, _ := http.NewRequest("POST", "https://oauth.battle.net/token", strings.NewReader(formData.Encode()))
	req.SetBasicAuth(client_id, client_secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Error occurred %s", err)
		return nil
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Recieved bad status code: %d", resp.StatusCode)
		return nil
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	if body != nil {
		var responseData AuthResponse
		err := json.Unmarshal(body, &responseData)
		if err != nil {
			fmt.Printf("Unmarshal error %s", err)
			return nil
		}
		return &responseData
	}
	return nil
}
