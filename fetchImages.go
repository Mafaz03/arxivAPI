package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func getImages(searchQuery string) (string, error) {

	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("bing api key not found")

	}
	bingapi := os.Getenv("BINGAPI")
	const endpoint = "https://api.bing.microsoft.com/v7.0/images/search"

	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return "", fmt.Errorf("unable to perform http request")
	}
	q := req.URL.Query()
	q.Add("q", searchQuery)
	q.Add("count", "1")
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Ocp-Apim-Subscription-Key", bingapi)
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	// Parse the JSON response
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("error parsing JSON: %v", err)
	}
	if value, found := result["value"]; found {
		if images, ok := value.([]interface{}); ok && len(images) > 0 {
			if firstImage, ok := images[0].(map[string]interface{}); ok {
				if contentUrl, found := firstImage["contentUrl"]; found {
					if url, ok := contentUrl.(string); ok {
						return url, nil
					}
					return "", fmt.Errorf("contentUrl is not a string")
				}
				return "", fmt.Errorf("contentUrl not found in the first image result")
			}
			return "", fmt.Errorf("first image result is not a valid map structure")
		}
		return "", fmt.Errorf("no images found or 'value' is not a valid array")
	}
	return "", fmt.Errorf("no 'value' key found in the response")
}
