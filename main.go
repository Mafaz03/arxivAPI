package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"github.com/Mafaz03/arxivAPI/internal/arxivapi"
)



func main() {

	client := arxivapi.NewClient(time.Minute)
	xmlData := client.FetchPapers()

	x := arxivapi.Feed{}

	err := xml.Unmarshal([]byte(xmlData), &x)
	if err != nil {
		log.Fatal(err)
	}
	if len(x.Entry) == 0 {
		fmt.Println("Empty")
		return
	}
	for i := range len(x.Entry) {
		fmt.Println(x.Entry[i].Title + "\n")
	}
}
