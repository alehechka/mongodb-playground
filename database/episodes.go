package database

import (
	"context"

	"github.com/alehechka/mongodb-playground/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertPodcastEpisode inserts a single episode record into the episode collection
func InsertPodcastEpisode(ctx context.Context, podcastID primitive.ObjectID, episode types.Episode) (primitive.ObjectID, error) {
	_, err := FindPodcast(ctx, types.Podcast{ID: podcastID})
	if err != nil {
		return primitive.NilObjectID, err
	}

	episode.ID = primitive.NewObjectID()
	episode.PodcastID = podcastID

	res, err := episodeCollection.InsertOne(ctx, episode)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

// FindPodcastEpisodes returns all episode records that match the provided filters
func FindPodcastEpisodes(ctx context.Context, podcastID primitive.ObjectID, find types.Episode) (episodes types.Episodes, err error) {
	find.PodcastID = podcastID
	cursor, err := episodeCollection.Find(ctx, find)
	if err != nil {
		return episodes, err
	}

	err = cursor.All(ctx, &episodes)
	episodes.Init()
	return
}

// FindPodcastEpisode returns a single episode that matches the provided parameters
func FindPodcastEpisode(ctx context.Context, podcastID primitive.ObjectID, find types.Episode) (episode types.Episode, err error) {
	find.PodcastID = podcastID
	res := episodeCollection.FindOne(ctx, find)
	if res.Err() != nil {
		return episode, res.Err()
	}

	err = res.Decode(&episode)
	return
}

// ReplacePodcastEpisode replaces the document with given ID with the data of given Episode
func ReplacePodcastEpisode(ctx context.Context, podcastID primitive.ObjectID, id primitive.ObjectID, update types.Episode) (err error) {
	update.ID = id
	update.PodcastID = podcastID

	res := episodeCollection.FindOneAndReplace(ctx, types.Episode{ID: id, PodcastID: podcastID}, update)
	return res.Err()
}

// DeletePodcastEpisode deletes the episode document with given ID
func DeletePodcastEpisode(ctx context.Context, podcastID primitive.ObjectID, id primitive.ObjectID) (err error) {
	res := episodeCollection.FindOneAndDelete(ctx, types.Episode{ID: id, PodcastID: podcastID})
	return res.Err()
}

// DeletePodcastEpisodes deletes all episodes of a podcast
func DeletePodcastEpisodes(ctx context.Context, podcastID primitive.ObjectID) (deleted int64, err error) {
	res, err := episodeCollection.DeleteMany(ctx, types.Episode{PodcastID: podcastID})

	return res.DeletedCount, err
}
