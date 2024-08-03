package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"apis"
)

const url = "http://export.arxiv.org/api/query?search_query=all:ass&start=0&max_results=2&sortBy=submittedDate&sortOrder=descending"

type Client struct {
	httpClient http.Client
}

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


func main() {

	client := &Client{
		httpClient: http.Client{},
	}
	xmlData := client.FetchPapers()

	x := arxivStruct{}

	err := xml.Unmarshal([]byte(xmlData), &x)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(x.Entries)
}
