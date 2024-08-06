package main

import (
	// "encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Mafaz03/arxivAPI/ai"
	"github.com/Mafaz03/arxivAPI/internal/arxivapi"
	// "github.com/Mafaz03/arxivAPI/ai"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	// "go.mongodb.org/mongo-driver/bson"
)

func fetch(amount int, db string, collection string) arxivapi.Feed{
	
	client := arxivapi.NewClient(time.Minute)
	xmlData := client.FetchPapers(amount, db, collection)

	x := arxivapi.Feed{}
	err := xml.Unmarshal([]byte(xmlData), &x)

	for i := 0; i < len(x.Entry); i++ {
        x.Entry[i].NewsTitle = ai.GetTitle(x.Entry[i].Title)
    }
	if err != nil {
		log.Fatal(err)
	}
	if len(x.Entry) == 0 {
		fmt.Println("Fetched Data returned empty")
		return arxivapi.Feed{}
	}

	return x
}


func main() {
	
	// content := ai.GetCompletion()
	// fmt.Println(content)
	
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r1router := chi.NewRouter()
	r1router.Get("/healthz", handler_readiness)
	r1router.Get("/getInfo/{db_col_amount}", get_Feeds)

	router.Mount("/v1", r1router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + "8080",
	}
	log.Printf("listening on Port number: %v", 8080)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("ERROR: ", err)
	}

	// feed := worker.fetchData()
	
}
