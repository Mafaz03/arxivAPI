package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Mafaz03/arxivAPI/internal/arxivapi"
	"github.com/Mafaz03/arxivAPI/internal/timeInfo"
)


func main() {
	timeinfo.readData
	data, err := timeinfo.readData("timeInfo.json")
	if err != nil {
        fmt.Println(err)
    }
    fmt.Println(data.Last_run_time)
	data.Last_run_time = time.Now().Format("2006-01-02 15:04:05")
	writeData("timeInfo.json", data)
	fmt.Println(data.Last_run_time)

	client := arxivapi.NewClient(time.Minute)
	xmlData := client.FetchPapers()

	x := arxivapi.Feed{}

	err = xml.Unmarshal([]byte(xmlData), &x)
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
}
