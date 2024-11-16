package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Conn *mongo.Database
}

func ConfigMongo(uri string, databaseName string) (*Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		return nil, err
	}

	return &Database{
		Conn: client.Database(databaseName),
	}, nil
}
