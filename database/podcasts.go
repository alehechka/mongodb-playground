package database

import (
	"fmt"

	"github.com/alehechka/mongodb-playground/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertPodcast inserts a single podcast record into the podcast collection
func InsertPodcast(podcast types.Podcast) (primitive.ObjectID, error) {
	podcast.ID = primitive.NewObjectID()
	fmt.Println(podcast.ID)

	res, err := podcastCollection.InsertOne(timeoutContext, podcast)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

// FindAllPodcasts returns all podcast records
func FindAllPodcasts() (podcasts types.Podcasts, err error) {
	cursor, err := podcastCollection.Find(timeoutContext, bson.M{})
	if err != nil {
		return podcasts, err
	}

	err = cursor.All(timeoutContext, &podcasts)
	return
}

// FindPodcasts returns all podcast records that match the provided filters
func FindPodcasts(find types.Podcast) (podcasts types.Podcasts, err error) {
	cursor, err := podcastCollection.Find(timeoutContext, find)
	if err != nil {
		return podcasts, err
	}

	err = cursor.All(timeoutContext, &podcasts)
	return
}

// FindPodcast returns a single podcast that matches the provided parameters
func FindPodcast(find types.Podcast) (podcast types.Podcast, err error) {
	res := podcastCollection.FindOne(timeoutContext, find)
	if res.Err() != nil {
		return podcast, res.Err()
	}

	err = res.Decode(&podcast)
	return
}

// ReplacePodcast replaces the document with given ID with the data of given Podcast
func ReplacePodcast(id primitive.ObjectID, update types.Podcast) (err error) {
	res := podcastCollection.FindOneAndReplace(timeoutContext, types.Podcast{ID: id}, update)
	return res.Err()
}

// DeletePodcast deletes the podcast document with given ID
func DeletePodcast(id primitive.ObjectID) (err error) {
	res := podcastCollection.FindOneAndDelete(timeoutContext, types.Podcast{ID: id})
	return res.Err()
}
