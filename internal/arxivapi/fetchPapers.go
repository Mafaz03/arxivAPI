package arxivapi

import (
	"io"
	"log"
	"net/http"
)

const url = "http://export.arxiv.org/api/query?search_query=cat:cs.AI&start=0&max_results=100&sortBy=submittedDate&sortOrder=descending"

func (c *Client) FetchPapers() string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}