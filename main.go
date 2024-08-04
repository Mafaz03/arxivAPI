package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"github.com/Mafaz03/arxivAPI/internal/arxivapi"
	"github.com/Mafaz03/arxivAPI/internal/timeinfo"
)

func fetch() {
	client := arxivapi.NewClient(time.Minute)
	xmlData := client.FetchPapers()

	x := arxivapi.Feed{}

	err := xml.Unmarshal([]byte(xmlData), &x)
	if err != nil {
		log.Fatal(err)
	}
	if len(x.Entry) == 0 {
		// fmt.Println("Empty")
		return
	}

	// for i := range len(x.Entry) {
	// 	// fmt.Println(x.Entry[i].Title + "\n")
	// }

	for i := range 5 {
		fmt.Println(x.Entry[i].Title + "\n\n")
	}
}

func main() {

	data, err := timeinfo.ReadData("timeInfo.json")
	if err != nil {
		fmt.Println(err)
	}

	now_time := time.Now().UTC()

	diff := now_time.Sub(data.LastRunTimeParsed)
	if diff > time.Hour {
		fmt.Println("Updating the times (its been over 1 hour(s) since last update)")
		err = timeinfo.UpdateLastRunTime("timeInfo.json")
		if err != nil {
			fmt.Println("Error updating last run time:", err)
			return
		}
	}

	fetch()

}
