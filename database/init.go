package database

import (
	"context"
	"time"

	"github.com/alehechka/mongodb-playground/constants"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var database *mongo.Database
var podcastCollection *mongo.Collection
var episodeCollection *mongo.Collection
var timeoutContext context.Context

// InitializeMongoDB initializes global MongoDB client
func InitializeMongoDB() (disconnect func() error, err error) {
	timeoutContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err = mongo.Connect(timeoutContext, options.Client().ApplyURI(constants.MongoDBConnectionString))
	if err != nil {
		cancel()
		return nil, err
	}

	if err := client.Ping(timeoutContext, readpref.Primary()); err != nil {
		cancel()
		return nil, err
	}

	database = client.Database(constants.DatabaseName)
	podcastCollection = database.Collection(constants.PodcastCollectionName)
	episodeCollection = database.Collection(constants.EpisodeCollectionName)

	return func() error {
		err := client.Disconnect(timeoutContext)
		cancel()
		return err
	}, nil
}
