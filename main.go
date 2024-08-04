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
)

type timeInfo struct {
	Static_info string `json:"static_info"`
	Last_run_time string `json:"last_run_time"`
}

func readData(filename string) (*timeInfo, error) {
    encData, err := os.ReadFile(filename)
    if err != nil {
        return nil, errors.New("time related json file could not be read")
    }

    timeinfo := &timeInfo{}
    err = json.Unmarshal(encData, timeinfo)
    if err != nil {
        return nil, errors.New("error unmarshaling json data")
    }

    return timeinfo, nil
}

func main() {
	data, _ := readData("timeInfo.json")
	
    fmt.Println(data.Last_run_time)

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
}
