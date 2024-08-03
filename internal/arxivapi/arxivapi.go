package arxivapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewClient(interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: interval,
		},
		
	}
}