package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Mafaz03/arxivAPI/internal/arxivapi"
	"github.com/joho/godotenv"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoServer struct {
	client *mongo.Client
}

func newMongoServer() *mongoServer {
	if err := godotenv.Load(); err != nil {
		log.Fatal("add correct mongoDB connectio string in envirnment")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	mongo_client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("could not connect to MongoDB, make sure the instance is valid and running")
	}

	return &mongoServer{
		client: mongo_client,
	}
}

func (worker *mongoServer) addData(doc arxivapi.Feed) {
	coll := worker.client.Database("cs").Collection("AI")
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

}

type singleEntry struct {
	Updated   string `json:"updated"`
	Published string `json:"published"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Author    []struct {
		Name string `json:"name"`
	} `json:"author"`
}

func (worker *mongoServer) fetchData() map[int]singleEntry {
	coll := worker.client.Database("cs").Collection("AI")

	var result arxivapi.Feed
	findOptions := options.FindOne().SetSort(bson.D{{Key: "updated", Value: -1}})

	// Delete the oldest if count excedes
	count, err := coll.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if count > 5 {
		fmt.Printf("Documents in Database (%v) has execed the limit (5), continuing to Delete the oldest", count)
		findOptions_Delete := options.FindOneAndDelete().SetSort(bson.D{{Key: "updates", Value: 1}})

		err := coll.FindOneAndDelete(context.TODO(), bson.D{}, findOptions_Delete)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = coll.FindOne(context.TODO(), bson.D{}, findOptions).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	jsonData := make(map[int]singleEntry)

	for i, entry := range result.Entry {
		// Create a new entry for Feedjson
		newEntry := struct {
			Updated   string `json:"updated"`
			Published string `json:"published"`
			Title     string `json:"title"`
			Summary   string `json:"summary"`
			Author    []struct {
				Name string `json:"name"`
			} `json:"author"`
		}{
			Updated:   entry.Updated,
			Published: entry.Published,
			Title:     entry.Title,
			Summary:   entry.Summary,
			Author: []struct {
				Name string "json:\"name\""
			}(entry.Author),
		}

		jsonData[i] = newEntry
	}

	return jsonData
}
