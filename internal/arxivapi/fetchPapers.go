package arxivapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
)


func (c *Client) FetchPapers(amount int, db string, collection string) string {
	url := fmt.Sprintf("http://export.arxiv.org/api/query?search_query=cat:%s.%s&start=0&max_results=%v&sortBy=submittedDate&sortOrder=descending", db, collection, amount)
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