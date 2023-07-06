package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	apiKey = "L08NHLNWUcgQrH0qRrawHg==9owwxmyF4HdaWA2l"
	apiUrl = "https://api.api-ninjas.com/v1/quotes"
)

type quotes struct {
	data []motto
}
type motto struct {
	Quote    string `json:"quote"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

func FetchMotto() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return ""
	}

	req.Header.Set("X-Api-Key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return ""
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return ""
	}

	var data []motto
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Failed to parse response to JSON", err)
		return ""
	}

	return fmt.Sprintf("[%s]\n%s", data[0].Author, data[0].Quote)
}
