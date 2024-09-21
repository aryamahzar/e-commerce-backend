package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// InitMongoDB initializes the MongoDB connection
func InitMongoDB() *mongo.Client {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	log.Println("Connected to MongoDB!")
	return client
}

// GetCollection returns a MongoDB collection from the "ecommerce" database
func GetCollection(collectionName string) *mongo.Collection {
	return client.Database("ecommerce").Collection(collectionName)
}

// CloseMongoDB closes the MongoDB connection
func CloseMongoDB() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal("Error while disconnecting MongoDB:", err)
	} else {
		log.Println("MongoDB connection closed.")
	}
}
