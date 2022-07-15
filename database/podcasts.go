package database

import (
	"sync"

	"github.com/alehechka/mongodb-playground/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertPodcast inserts a single podcast record into the podcast collection
func InsertPodcast(podcast types.Podcast) (primitive.ObjectID, error) {
	podcast.ID = primitive.NewObjectID()

	res, err := podcastCollection.InsertOne(timeoutContext, podcast)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

// FindPodcasts returns all podcast records that match the provided filters
func FindPodcasts(find types.Podcast) (podcasts types.Podcasts, err error) {
	cursor, err := podcastCollection.Find(timeoutContext, find)
	if err != nil {
		return podcasts, err
	}

	err = cursor.All(timeoutContext, &podcasts)
	podcasts.Init()
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

// GetPodcastWithEpisodes retrieves a single podcast with all episodes included
func GetPodcastWithEpisodes(podcastID primitive.ObjectID) (podcast types.Podcast, err error) {
	var wg sync.WaitGroup

	wg.Add(1)
	var pErr error
	go func() {
		podcast, pErr = FindPodcast(types.Podcast{ID: podcastID})
		wg.Done()
	}()

	wg.Add(1)
	var episodes types.Episodes
	var eErr error
	go func() {
		episodes, eErr = FindPodcastEpisodes(podcastID, types.Episode{})
		wg.Done()
	}()

	wg.Wait()
	if pErr != nil {
		return types.Podcast{}, pErr
	}
	if eErr != nil {
		return types.Podcast{}, eErr
	}

	podcast.Episodes = episodes
	return
}

// ReplacePodcast replaces the document with given ID with the data of given Podcast
func ReplacePodcast(id primitive.ObjectID, update types.Podcast) (err error) {
	update.ID = id
	res := podcastCollection.FindOneAndReplace(timeoutContext, types.Podcast{ID: id}, update)
	return res.Err()
}

// DeletePodcast deletes the podcast document with given ID
func DeletePodcast(id primitive.ObjectID) (err error) {
	res := podcastCollection.FindOneAndDelete(timeoutContext, types.Podcast{ID: id})
	if res.Err() != nil {
		return res.Err()
	}

	_, err = DeletePodcastEpisodes(id)

	return
}
