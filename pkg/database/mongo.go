package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConnectionInfo struct {
	URI      string `mapstructure:"uri"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func NewMongoConnection(ctx context.Context, info ConnectionInfo) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	opts := options.Client().
		ApplyURI(info.URI).
		SetAuth(options.Credential{
			Username: info.Username,
			Password: info.Password,
		})

	log.Println("Connecting to Mongo...")
	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err := dbClient.Ping(ctx, nil); err != nil {
		return nil, err
	}

	db := dbClient.Database(info.Database)
	log.Println("Mongo connection is ready.")

	return db, nil
}
