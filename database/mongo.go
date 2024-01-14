package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectToMongo(mongoUrl string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoUrl)

	c, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Panic("Error connecting:", err)
		return nil, err
	}

	log.Printf("Connected to MongoDB at %s\n", mongoUrl)

	return c, nil
}
