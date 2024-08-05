package main

import (
	// "encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"github.com/Mafaz03/arxivAPI/internal/arxivapi"
	"github.com/Mafaz03/arxivAPI/internal/timeinfo"
	// "go.mongodb.org/mongo-driver/bson"
)

func fetch() arxivapi.Feed{
	
	client := arxivapi.NewClient(time.Minute)
	xmlData := client.FetchPapers()

	x := arxivapi.Feed{}

	err := xml.Unmarshal([]byte(xmlData), &x)
	if err != nil {
		log.Fatal(err)
	}
	if len(x.Entry) == 0 {
		fmt.Println("Fetched Data returned empty")
		return arxivapi.Feed{}
	}

	for i := range 5 {
		fmt.Println(x.Entry[i].Title + "\n\n")
	}
	return x
}


func main() {

	
	data, err := timeinfo.ReadData("timeInfo.json")
	if err != nil {
		fmt.Println(err)
	}

	now_time := time.Now().UTC()

	worker := newMongoServer()

	diff := now_time.Sub(data.LastRunTimeParsed)
	if diff > time.Hour {
		fmt.Println("Updating the time, it has been over 1 hour(s) since last update")
		err = timeinfo.UpdateLastRunTime("timeInfo.json")
		if err != nil {
			fmt.Println("Error updating last run time:", err)
			return
		} 
		feed := fetch()
		worker.addData(feed)
	}

	feed := worker.fetchData()
	fmt.Println(feed)
	
}
