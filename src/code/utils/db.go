package dbutils

import (
	"context"
	"fmt"
	"log"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

var Connect = func() *mongo.Client {

	// Setting client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, conerr := mongo.Connect(context.TODO(), clientOptions)

	if conerr != nil {
		log.Fatal(conerr)
	}

	// Check the connection
	conerr = client.Ping(context.TODO(), nil)

	if conerr != nil {
		log.Fatal(conerr)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}
