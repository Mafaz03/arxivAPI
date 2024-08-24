package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Mafaz03/arxivAPI/internal/timeinfo"
	"github.com/go-chi/chi"
)

func handler_readiness(w http.ResponseWriter, r *http.Request) {

	type parameter struct {
		Message string `json:"Satus code [200]"`
	}
	params := parameter{
		Message: "You are ready to go!",
	}

	jsonData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&jsonData)
	if err != nil {
		log.Printf("There was an error decoding the params")
	}

	respondWithJSON(w, 200, params)
}

func get_Feeds(w http.ResponseWriter, r *http.Request) {
	data, err := timeinfo.ReadData("timeInfo.json")
	if err != nil {
		fmt.Println(err)
	}

	now_time := time.Now().UTC()

	worker := newMongoServer()
	db_col_amount := chi.URLParam(r, "db_col_amount")
	split1 := strings.Split(db_col_amount, "~")
	amount, err := strconv.Atoi(split1[1])
	if err != nil {
		log.Fatal("could not convert amount to int")
	}

	db_col := strings.Split(split1[0], ":")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	torf := checkDatabaseCollectionExists(ctx, worker.client, db_col[0], db_col[1])
	if !torf {
		feed := fetch(amount, db_col[0], db_col[1])
		worker.addData(feed, db_col[0], db_col[1])
	}

	feed, count := worker.fetchData(amount, db_col[0], db_col[1])
	diff := now_time.Sub(data.LastRunTimeParsed)
	fmt.Println(count)

	

	if diff > time.Hour || count == 0 {
		fmt.Println("Updating the time, it has been over 1 hour(s) since last update")
		err = timeinfo.UpdateLastRunTime("timeInfo.json")
		if err != nil {
			fmt.Println("Error updating last run time:", err)
			return
		} 
		feed := fetch(amount, db_col[0], db_col[1])
		worker.addData(feed, db_col[0], db_col[1])
	}

	
	
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&feed)
	if err != nil {
		log.Printf("There was an error decoding the params: %v", err)
	}
	respondWithJSON(w, 200, feed)
	
}
