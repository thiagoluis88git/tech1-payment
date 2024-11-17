package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	Conn *mongo.Database
}

func ConfigMongo(client *mongo.Client, databaseName string) (*Database, error) {
	return &Database{
		Conn: client.Database(databaseName),
	}, nil
}
