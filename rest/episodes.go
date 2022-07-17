package rest

import (
	"net/http"

	"github.com/alehechka/go-utils/ginshared"
	"github.com/alehechka/mongodb-playground/database"
	"github.com/alehechka/mongodb-playground/opentel"
	"github.com/alehechka/mongodb-playground/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.opentelemetry.io/otel/attribute"
)

func getPodcastEpisodes(c *gin.Context) {
	ctx, span := opentel.GinTracer.Start(c.Request.Context(), "getPodcastEpisodes")
	defer span.End()

	podcastID, err := primitive.ObjectIDFromHex(c.Param(types.PodcastID))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		span.RecordError(err)
		return
	}
	span.SetAttributes(attribute.String(types.PodcastID, podcastID.Hex()))

	filterEpisode := types.Episode{
		Title:       c.Query("filter[title]"),
		Description: c.Query("filter[description]"),
	}
	span.SetAttributes(attribute.String("filterEpisode", filterEpisode.String()))

	episodes, err := database.FindPodcastEpisodes(ctx, podcastID, filterEpisode)
	if ginshared.ShouldAbortWithError(c)(http.StatusInternalServerError, err) {
		span.RecordError(err)
		return
	}

	c.JSON(http.StatusOK, types.EpisodesResponse{Episodes: episodes})
}

func getPodcastEpisode(c *gin.Context) {
	ctx, span := opentel.GinTracer.Start(c.Request.Context(), "getPodcastEpisode")
	defer span.End()

	podcastID, err := primitive.ObjectIDFromHex(c.Param(types.PodcastID))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		span.RecordError(err)
		return
	}
	span.SetAttributes(attribute.String(types.PodcastID, podcastID.Hex()))

	episodeID, err := primitive.ObjectIDFromHex(c.Param(types.EpisodeID))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		span.RecordError(err)
		return
	}
	span.SetAttributes(attribute.String(types.EpisodeID, episodeID.Hex()))

	episode, err := database.FindPodcastEpisode(ctx, podcastID, types.Episode{ID: episodeID})
	if ginshared.ShouldAbortWithError(c)(http.StatusNotFound, err) {
		span.RecordError(err)
		return
	}

	c.JSON(http.StatusOK, types.EpisodeResponse{Episode: episode})
}

func createPodcastEpisode(c *gin.Context) {
	ctx, span := opentel.GinTracer.Start(c.Request.Context(), "createPodcastEpisode")
	defer span.End()

	podcastID, err := primitive.ObjectIDFromHex(c.Param(types.PodcastID))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		span.RecordError(err)
		return
	}
	span.SetAttributes(attribute.String(types.PodcastID, podcastID.Hex()))

	var episode types.Episode
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, c.ShouldBind(&episode)) {
		span.RecordError(err)
		return
	}
	span.SetAttributes(attribute.String("episode", episode.String()))

	episodeID, err := database.InsertPodcastEpisode(ctx, podcastID, episode)
	if ginshared.ShouldAbortWithError(c)(http.StatusInternalServerError, err) {
		span.RecordError(err)
		return
	}

	c.Params = append(c.Params, gin.Param{Key: types.EpisodeID, Value: episodeID.Hex()})
	getPodcastEpisode(c)
}

func replacePodcastEpisode(c *gin.Context) {
	ctx, span := opentel.GinTracer.Start(c.Request.Context(), "replacePodcastEpisode")
	defer span.End()

	podcastID, err := primitive.ObjectIDFromHex(c.Param(types.PodcastID))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		span.RecordError(err)
		return
	}
	span.SetAttributes(attribute.String(types.PodcastID, podcastID.Hex()))

	episodeID, err := primitive.ObjectIDFromHex(c.Param(types.EpisodeID))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		span.RecordError(err)
		return
	}
	span.SetAttributes(attribute.String(types.EpisodeID, episodeID.Hex()))

	var episode types.Episode
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, c.ShouldBind(&episode)) {
		span.RecordError(err)
		return
	}
	span.SetAttributes(attribute.String("episode", episode.String()))

	err = database.ReplacePodcastEpisode(ctx, podcastID, episodeID, episode)
	if ginshared.ShouldAbortWithError(c)(http.StatusInternalServerError, err) {
		span.RecordError(err)
		return
	}

	getPodcastEpisode(c)
}

func deletePodcastEpisode(c *gin.Context) {
	ctx, span := opentel.GinTracer.Start(c.Request.Context(), "deletePodcastEpisode")
	defer span.End()

	podcastID, err := primitive.ObjectIDFromHex(c.Param(types.PodcastID))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		span.RecordError(err)
		return
	}
	span.SetAttributes(attribute.String(types.PodcastID, podcastID.Hex()))

	episodeID, err := primitive.ObjectIDFromHex(c.Param(types.EpisodeID))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		span.RecordError(err)
		return
	}
	span.SetAttributes(attribute.String(types.EpisodeID, episodeID.Hex()))

	err = database.DeletePodcastEpisode(ctx, podcastID, episodeID)
	if ginshared.ShouldAbortWithError(c)(http.StatusNotFound, err) {
		span.RecordError(err)
		return
	}

	c.Status(http.StatusNoContent)
}
