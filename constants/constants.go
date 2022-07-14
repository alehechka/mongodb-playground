package constants

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// MongoDBConnectionString represents the uri to connect to a MongoDB instance
var MongoDBConnectionString string

// MongoDB database and collection constants
const (
	// DatabaseName is the name of the MongoDB database
	DatabaseName = "streaming-schema"

	// PodcastCollectionName is the name of the podcast collection
	PodcastCollectionName = "podcasts"

	// EpisodeCollectionName is the name of the podcast collection
	EpisodeCollectionName = "episodes"
)

//InitializeConstants initializes global constant values from env
func InitializeConstants() {
	err := godotenv.Load("secrets/.env")
	if err != nil {
		fmt.Println(err.Error())
	}

	MongoDBConnectionString = os.Getenv("MONGODB_CONNECTION_STRING")
}
