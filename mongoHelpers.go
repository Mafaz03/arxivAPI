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

func (worker *mongoServer) addData(doc arxivapi.Feed, db string, collection string) {
	coll := worker.client.Database(db).Collection(collection)
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
	NewsTitle string `json:"newstitle"`
	Summary   string `json:"summary"`
	Author    []struct {
		Name string `json:"name"`
	} `json:"author"`
}

func checkDatabaseExists(ctx context.Context, client *mongo.Client, databaseName string) (bool, error) {
	databases, err := client.ListDatabaseNames(ctx, nil)
	if err != nil {
		return false, err
	}

	for _, db := range databases {
		if db == databaseName {
			return true, nil
		}
	}
	return false, nil
}

func checkCollectionExists(ctx context.Context, client *mongo.Client, databaseName string, collectionName string) (bool, error) {
	collections, err := client.Database(databaseName).ListCollectionNames(ctx, nil)
	if err != nil {
		return false, err
	}

	for _, coll := range collections {
		if coll == collectionName {
			return true, nil
		}
	}
	return false, nil
}

func checkDatabaseCollectionExists(ctx context.Context, client *mongo.Client, databaseName string, collectionName string) bool {
	// Check if the database exists
	databaseExists, err := checkDatabaseExists(ctx, client, databaseName)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if !databaseExists {
		return false
	}

	// Check if the collection exists
	collectionExists, err := checkCollectionExists(ctx, client, databaseName, collectionName)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return collectionExists
}

func (worker *mongoServer) fetchData(amount int, db string, collection string) (map[int]singleEntry, int64) {
	coll := worker.client.Database(db).Collection(collection)

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
			NewsTitle string `json:"newstitle"`
			Summary   string `json:"summary"`
			Author    []struct {
				Name string `json:"name"`
			} `json:"author"`
		}{
			Updated:   entry.Updated,
			Published: entry.Published,
			Title:     entry.Title,
			NewsTitle: entry.NewsTitle,
			Summary:   entry.Summary,
			Author: []struct {
				Name string "json:\"name\""
			}(entry.Author),
		}

		jsonData[i] = newEntry
		if i+1 >= amount {
			break
		}
	}

	return jsonData, count
}
