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

func (worker *mongoServer) fetchData() {
	coll := worker.client.Database("cs").Collection("AI")
	
	var result arxivapi.Feed
	findOptions := options.FindOne().SetSort(bson.D{{Key: "updated", Value: -1}})

	err := coll.FindOne(context.TODO(), bson.D{}, findOptions).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf(result.Entry[0].Summary)

}
