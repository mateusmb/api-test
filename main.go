package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var username = "blackeagle"
var password = "n1ghtlov3"
var host1 = "cluster0-8tjrn.gcp.mongodb.net/test"

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	ctx := context.TODO()

	//pw, ok := os.LookupEnv("MONGO_PW")
	//
	//if !ok {
	//	fmt.Println("error: Unable to find MONGO_PW in the environment")
	//	os.Exit(1)
	//}

	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s", username, password, host1)
	fmt.Println("connection string is: ", mongoURI)

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Connection error")
		log.Fatal(err)
		os.Exit(1)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		fmt.Println("Ping error")
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("people")

	ash := Person{"Ash", 10, "Pallet Town"}
	misty := Person{"Misty", 10, "Cerulean City"}
	brock := Person{"Brock", 15, "Pewter City"}

	inserResult, err := collection.InsertOne(context.TODO(), ash)

	if err != nil {
		fmt.Println("Insertion error")
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", inserResult.InsertedID)

	people := []interface{}{misty, brock}

	insertManyResult, err := collection.InsertMany(context.TODO(), people)

	if err != nil {
		fmt.Println("Insert many error")
		log.Fatal(err)
	}

	fmt.Println("Inserted many documents: ", insertManyResult.InsertedIDs)
}
