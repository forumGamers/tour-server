package config

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

func Connection() {
	if err := godotenv.Load() ; err != nil {
		panic(err.Error())
	}

	uri := os.Getenv("URI")

	if uri == "" {
		uri = "mongodb://localhost:27017"
	}
	options := options.Client().ApplyURI(uri)
	if client,err := mongo.Connect(context.Background(),options) ; err != nil {
		panic(err.Error())
	}else {
		Db = client.Database("Tour")
		fmt.Println("connection success")
	}
}