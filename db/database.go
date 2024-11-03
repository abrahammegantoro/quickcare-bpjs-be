package db

import (
	"context"
	"fmt"

	"github.com/abrahammegantoro/quickcare-bpjs-be/config"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Init(config *config.EnvConfig) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.DBUrl)
	
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: %e", err)
		return nil, err
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Error pinging MongoDB: %e", err)
		return nil, err
	}

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Connected to MongoDB!")
	return client, nil
}
