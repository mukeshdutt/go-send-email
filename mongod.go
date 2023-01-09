package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// database details
const APPDB = "e-mailer"
const CONNECTION = "mongodb://localhost:27017"

// Created a global variable which holds the connection details
var mongoClient *mongo.Client

// created the database instance on the pakcage initialization
func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(CONNECTION))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal()
	}
	mongoClient = client
}
