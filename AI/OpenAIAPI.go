package ai

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

func getCompletion() string{
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
			"messages":   []interface{}{map[string]interface{}{"role": "system", "content": "Hi can you tell me what is the factorial of 10?"}},
			"max_tokens": 50,
		}).Post(apiEndpoint)

	if err != nil {
		log.Fatalf("Error while sending send the request: %v", err)
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(response.Body(), &data)
	if err != nil {
        fmt.Println("Error while decoding JSON response:", err)
        return ""
    }
	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	return content
}
