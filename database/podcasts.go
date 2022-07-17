package database

import (
	"context"
	"sync"

	"github.com/alehechka/mongodb-playground/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertPodcast inserts a single podcast record into the podcast collection
func InsertPodcast(ctx context.Context, podcast types.Podcast) (primitive.ObjectID, error) {
	podcast.ID = primitive.NewObjectID()

	res, err := podcastCollection.InsertOne(ctx, podcast)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

// FindPodcasts returns all podcast records that match the provided filters
func FindPodcasts(ctx context.Context, find types.Podcast) (podcasts types.Podcasts, err error) {
	cursor, err := podcastCollection.Find(ctx, find, &options.FindOptions{})
	if err != nil {
		return podcasts, err
	}

	err = cursor.All(ctx, &podcasts)
	podcasts.Init()

	return
}

// FindPodcast returns a single podcast that matches the provided parameters
func FindPodcast(ctx context.Context, find types.Podcast) (podcast types.Podcast, err error) {
	res := podcastCollection.FindOne(ctx, find)
	if res.Err() != nil {
		return podcast, res.Err()
	}

	err = res.Decode(&podcast)
	return
}

// GetPodcastWithEpisodes retrieves a single podcast with all episodes included
func GetPodcastWithEpisodes(ctx context.Context, podcastID primitive.ObjectID) (podcast types.Podcast, err error) {
	var wg sync.WaitGroup

	wg.Add(1)
	var pErr error
	go func() {
		podcast, pErr = FindPodcast(ctx, types.Podcast{ID: podcastID})
		wg.Done()
	}()

	wg.Add(1)
	var episodes types.Episodes
	var eErr error
	go func() {
		episodes, eErr = FindPodcastEpisodes(ctx, podcastID, types.Episode{})
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
func ReplacePodcast(ctx context.Context, id primitive.ObjectID, update types.Podcast) (err error) {
	update.ID = id
	res := podcastCollection.FindOneAndReplace(ctx, types.Podcast{ID: id}, update)
	return res.Err()
}

// DeletePodcast deletes the podcast document with given ID
func DeletePodcast(ctx context.Context, id primitive.ObjectID) (err error) {
	res := podcastCollection.FindOneAndDelete(ctx, types.Podcast{ID: id})
	if res.Err() != nil {
		return res.Err()
	}

	_, err = DeletePodcastEpisodes(ctx, id)

	return
}
