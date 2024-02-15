package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func NewMongoClient() *mongo.Client {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017/").SetWriteConcern(writeconcern.W1()).SetReadConcern(readconcern.Local())
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	return client
}
