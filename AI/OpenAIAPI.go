package ai

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

func GetTitle(title string) (string, string){
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	openaiapi := os.Getenv("OPENAIAPI")
	if openaiapi == "" {
		log.Fatal("enter valid api key, learn more at: https://platform.openai.com/api-keys")
	}
	
	client := resty.New()
	response, err := client.R().SetAuthToken(openaiapi).SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model":      "gpt-3.5-turbo",
			"messages":   []interface{}{map[string]interface{}{"role": "system", "content": fmt.Sprintf(
				`Generate news title on: %s, 
				last line must be related to title content from which an amazing image can be fetched, the image search query must be very simple

				example 1:
				Investigating Emergence in Transformers Trained on Formal Language through a Percolation Model
				Percolation Model amazing pictures

				example 2:
				Unveiling the Power of Line Charts: A Doorway to Dataset Exploration and Insight
				cool images on line charts
				`,
				title)}},
			"max_tokens": 50,
		}).Post(apiEndpoint)

	if err != nil {
		log.Fatalf("Error while sending send the request: %s", err)
	}
	var data map[string]interface{}
	err = json.Unmarshal(response.Body(), &data)
	if err != nil {
        fmt.Println("Error while decoding JSON response:", err)
        return "", ""
    }
	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	genText := strings.Split(content, "\n")

	genTitle := genText[0]
	imageQuery := genText[1]
	return genTitle, imageQuery
}
